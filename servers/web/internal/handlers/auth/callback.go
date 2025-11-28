package auth

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"zotes/servers/web/db"
	"zotes/servers/web/internal/app"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserInfo struct {
	Name     string
	Nickname string
	Picture  string
	ID       string
}

type CallbackProps struct {
	r *http.Request
}

func (crp *CallbackProps) ParseParams() (CallbackRequestParams, error) {
	error := crp.r.URL.Query().Get("error")
	if error != "" {
		return CallbackRequestParams{}, fmt.Errorf("CallbackRequestParams(ParseParams): %s", error)
	}
	state := crp.r.URL.Query().Get("state")
	code := crp.r.URL.Query().Get("code")
	if state == "" {
		return CallbackRequestParams{}, fmt.Errorf("CallbackRequestParams(ParseParams): missing state")
	}
	if code == "" {
		return CallbackRequestParams{}, fmt.Errorf("CallbackRequestParams(ParseParams): missing code")
	}
	return CallbackRequestParams{
		State: state,
		Code:  code,
	}, nil
}

type CallbackRequestParams struct {
	State string
	Code  string
}

func (crpa *CallbackRequestParams) ValidateID(r context.Context, tools *app.App) (*oidc.IDToken, error) {
	storedState := tools.Session.Get(r, "state")
	if storedState == "" {
		return &oidc.IDToken{}, fmt.Errorf("CallbackRequestParams(ValidateAndGetToken): missing state")
	} else if storedState != crpa.State {
		tools.Logger.Error("State mismatch", slog.Any("storedState", storedState), slog.Any("crpa.State", crpa.State))
		return &oidc.IDToken{}, fmt.Errorf("CallbackRequestParams(ValidateAndGetToken): state mismatch")
	}

	token, err := tools.OidcAuthenticator.Exchange(r, crpa.Code)
	if err != nil {
		return &oidc.IDToken{}, fmt.Errorf("CallbackRequestParams(ValidateAndGetToken)(Oidc.Exchange): %s", err)
	}

	idtoken, err := tools.OidcAuthenticator.VerifyIDToken(r, token)
	if err != nil {
		return &oidc.IDToken{}, fmt.Errorf("CallbackRequestParams(ValidateAndGetToken)(Oidc.Verify): %s", err)
	}
	tools.Session.Remove(r, "state")
	return idtoken, nil
}

type OidcProfile struct {
	Sub      string `json:"sub"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Picture  string `json:"picture"`
	Verified bool   `json:"email_verified"`
}

func GetProfile(idtoken *oidc.IDToken) (*OidcProfile, error) {
	var prof OidcProfile

	if err := idtoken.Claims(&prof); err != nil {
		return nil, fmt.Errorf("GetProfile: invalid claims: %w", err)
	}

	if prof.Sub == "" {
		return nil, fmt.Errorf("GetProfile: missing sub")
	}
	if prof.Name == "" {
		return nil, fmt.Errorf("GetProfile: missing name")
	}
	if prof.Email == "" {
		return nil, fmt.Errorf("GetProfile: missing email")
	}
	if !prof.Verified {
		return nil, fmt.Errorf("GetProfile: email not verified")
	}

	return &prof, nil
}

type NewUserSessionProps struct {
	Profile         *OidcProfile
	AuthenticatedID pgtype.UUID
}

func (nssp *NewUserSessionProps) GetAuthenticatedID() string {
	return nssp.AuthenticatedID.String()
}

func NewUserSession(
	ctx context.Context,
	props *NewUserSessionProps,
	tools *app.App,
) error {

	uid := props.GetAuthenticatedID()
	name := props.Profile.Name
	email := props.Profile.Email

	scstoken := ""

	err := tools.WithTx(ctx, app.RLSNone, func(q *db.Queries) error {
		// You can generate token here or outside
		tmpToken := tools.Session.Token(ctx)

		err := q.NewUserSession(ctx, db.NewUserSessionParams{
			UserID:    props.AuthenticatedID,
			LastToken: pgtype.Text{String: tmpToken, Valid: true},
		})
		if err != nil {
			return err
		}

		scstoken = tmpToken
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed user session tx: %w", err)
	}

	// ONLY NOW mutate the cookie/session.
	tools.Session.Put(ctx, "authenticatedUserID", uid)
	tools.Session.Put(ctx, "user_name", name)
	tools.Session.Put(ctx, "user_email", email)
	tools.Session.Put(ctx, "session_token", scstoken)

	return nil
}

func GetUserID(r context.Context, profile *OidcProfile, tools *app.App) (pgtype.UUID, error) {
	var id pgtype.UUID
	err := tools.WithTx(r, app.RLSNone, func(q *db.Queries) error {

		newID, err := q.UpsertUserBySub(r, pgtype.Text{
			String: profile.Sub,
			Valid:  true,
		})
		if err != nil {
			return err
		}

		id = newID
		return nil
	})
	if err != nil {
		return pgtype.UUID{}, fmt.Errorf("failed user session tx: %w", err)
	}

	return id, nil
}

func Callback(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		props := CallbackProps{
			r: r,
		}

		params, err := props.ParseParams()
		if err != nil {
			app.Logger.Error("Error response", slog.Any("error", err.Error()))
			errmsg := fmt.Sprintf("ParseParams: %s", err)
			http.Redirect(w, r, fmt.Sprintf("/auth/login?error=%s", url.QueryEscape(errmsg)), http.StatusSeeOther)
			return
		}

		idtoken, err := params.ValidateID(r.Context(), app)
		if err != nil {
			app.Logger.Error("User info incorrect", slog.Any("error", err.Error()))
			errmsg := fmt.Sprintf("ValidateID: %s", err)
			http.Redirect(w, r, fmt.Sprintf("/auth/login?error=%s", url.QueryEscape(errmsg)), http.StatusSeeOther)
			return
		}

		profile, err := GetProfile(idtoken)
		if err != nil {
			app.Logger.Error("User info incorrect", slog.Any("error", err.Error()))
			errmsg := fmt.Sprintf("GetProfile: %s", err)
			http.Redirect(w, r, fmt.Sprintf("/auth/login?error=%s", url.QueryEscape(errmsg)), http.StatusSeeOther)
			return
		}

		authenticatedUserID, err := GetUserID(r.Context(), profile, app)
		if err != nil {
			app.Logger.Error("GetUserID", slog.Any("error", err.Error()))
			errmsg := fmt.Sprintf("GetUserID: %s", err)
			http.Redirect(w, r, fmt.Sprintf("/auth/login?error=%s", url.QueryEscape(errmsg)), http.StatusSeeOther)
			return
		}
		//
		// group, err := GetGroup(r.Context(), authenticatedUserID, app)
		// if err != nil {
		// 	app.Logger().Error("GetGroupID", slog.Any("error", err.Error()))
		// 	http.Redirect(w, r, "/auth/login?error=true", http.StatusSeeOther)
		// 	return
		// }
		err = NewUserSession(r.Context(), &NewUserSessionProps{
			Profile:         profile,
			AuthenticatedID: authenticatedUserID,
			// Group:           group,
		}, app)
		if err != nil {
			app.Logger.Error("NewUserSession", slog.Any("error", err.Error()))
			http.Redirect(w, r, "/auth/login?error=true", http.StatusSeeOther)
			return
		}

		// err = updateDevSessionEnv(r, profile, app)
		// if err != nil {
		// 	app.Logger().Error("updateDevSessionEnv", slog.Any("error", err.Error()))
		// 	http.Redirect(w, r, "/auth/login?error=true", http.StatusSeeOther)
		// 	return
		// }
		w.Header().Set("Content-Type", "text/html")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// func updateDevSessionEnv(r *http.Request, profile *OidcProfile, tools *app.App) error {
// 	if os.Getenv("ENVIRONMENT") == "dev" {
// 		cookie := r.Cookies()
// 		var s string
// 		for _, c := range cookie {
// 			if c.Name == "session" {
// 				s = fmt.Sprintf("%s", c.Value)
// 			}
// 		}
//
// 		nickname := strings.ReplaceAll(profile.Name, "+", "_")
// 		nickname = strings.ReplaceAll(nickname, ".", "")
//
// 		tools.Logger.Info("updateDevSessionEnv", slog.Any("nickname", nickname), slog.Any("s", s))
//
// 		path := "http/.env"
// 		lines := []string{}
// 		key := fmt.Sprintf("%s=session", nickname)
// 		newEntry := fmt.Sprintf("%s=%s", key, s)
//
// 		// Read file if it exists
// 		if data, err := os.ReadFile(path); err == nil {
// 			for _, line := range strings.Split(string(data), "\n") {
// 				if !strings.HasPrefix(line, key+"=") && strings.TrimSpace(line) != "" {
// 					lines = append(lines, line)
// 				}
// 			}
// 		}
// 		lines = append(lines, newEntry)
//
// 		// Write updated content
// 		return os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), 0644)
// 	}
// 	return fmt.Errorf("updateDevSessionEnv: not in dev env")
// }
