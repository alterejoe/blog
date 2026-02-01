package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/alterejoe/blog/internal/app"

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

	return AuthenticatedUser{
		UserID: u,
		// GroupID: g,
	}, nil
}

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
