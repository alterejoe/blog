package main

import (
	"context"
	"fmt"
	"net/http"
	"zotes/servers/web/internal/app"

	"github.com/google/uuid"
)

type AuthenticatedUser struct {
	UserID uuid.UUID
	// GroupID uuid.UUID
}

func AuthPresentInSession(r context.Context, app *app.App) (AuthenticatedUser, error) {
	var user string
	sessionu := app.Session.Get(r, "authenticatedUserID")
	if sessionu == nil {
		return AuthenticatedUser{}, fmt.Errorf("user not in session")
	}
	user = sessionu.(string)

	u, err := uuid.Parse(user)
	if err != nil {
		return AuthenticatedUser{}, fmt.Errorf("invalid user uuid: %s", err)
	}

	// sessiong := app.Session().Get(r, "authenticatedClientGroup")
	// if sessiong == nil {
	// 	return AuthenticatedUser{}, fmt.Errorf("group not in session")
	// }
	//
	// group := sessiong.(string)
	//
	// g, err := uuid.Parse(group)
	// if err != nil {
	// 	return AuthenticatedUser{}, fmt.Errorf("invalid group uuid: %s", err)
	// }

	return AuthenticatedUser{
		UserID: u,
		// GroupID: g,
	}, nil
}

// func ValidateUserPresentInGroup(r context.Context, authuser AuthenticatedUser, tools *app.App) error {
// 	tx, cleanup, err := tools.TxSystemRLSWithCleanup(r)
// 	if err != nil {
// 		return fmt.Errorf("ValidateUserPresentInGroup(transaction creation): %s", err)
// 	}
// 	defer cleanup(err == nil)
//
// 	params := db.UserPresentInGroupParams{
// 		GroupID: pgtype.UUID{
// 			Bytes: authuser.GroupID,
// 			Valid: true,
// 		},
// 		UserID: pgtype.UUID{
// 			Bytes: authuser.UserID,
// 			Valid: true,
// 		},
// 	}
// 	var present bool
// 	p, err := tools.QueryTx(r, &sqlcqueries.UserPresentInGroup{
// 		Params: params,
// 	}, tx)
// 	if err != nil {
// 		return fmt.Errorf("ValidateUserPresentInGroup(database check): %s", err)
// 	}
//
// 	present, ok := p.(bool)
// 	if !ok {
// 		return fmt.Errorf("ValidateUserPresentInGroup(bool conversion): %s", err)
// 	}
// 	if !present {
// 		return fmt.Errorf("ValidateUserPresentInGroup(user not in group)")
// 	}
// 	return nil
// }

func ClientAuthMiddleware(app *app.App) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := AuthPresentInSession(r.Context(), app)
			if err != nil {
				app.Session.DeleteAuthUser(r.Context())
				http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
				return
			}

			// err = ValidateUserPresentInGroup(r.Context(), userauth, app)
			// if err != nil {
			// 	app.Session.DeleteAuthUser(r.Context())
			// 	http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			// 	return
			// }
			// if !IsAuthenticated(r.Context(), app) {
			// 	// app.Session().DeleteAuthUser(r.Context())
			// 	http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			// 	return
			// }

			w.Header().Add("Cache-Control", "no-store")
			next.ServeHTTP(w, r)
		})
	}
}
