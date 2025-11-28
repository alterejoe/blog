package auth

import (
	"log/slog"
	"net/http"
	"net/url"
	"zotes/servers/web/internal/app"
)

// TODO: OPS-10
func Login(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// session := app.GetRedisSessionManager()
		// auth := app.GetAuth()
		error := r.URL.Query().Get("error")
		if error != "" {
			app.Session.Destroy(r.Context())
			u := url.URL{
				Path: "/auth/landing",
			}
			q := u.Query()
			q.Add("error", error)
			u.RawQuery = q.Encode()
			http.Redirect(w, r, u.String(), http.StatusSeeOther)
			return
		}

		state, err := app.OidcAuthenticator.GenerateRandomState()
		// valid.CheckField(err == nil, "State", "Invalid State")
		if err != nil {
			app.Logger.Error("PostLogin", slog.Any("error", err.Error()))
			//500
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		app.Session.Put(r.Context(), "state", state)

		http.Redirect(w, r, app.OidcAuthenticator.AuthCodeURL(state), http.StatusSeeOther)
	}
}
