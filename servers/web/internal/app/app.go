package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"zotes/shared/interfaces"

	"github.com/casbin/casbin/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	DB                *pgxpool.Pool
	Logger            interfaces.Logger
	Session           interfaces.SessionManager
	OidcAuthenticator interfaces.OIDCAuthenticator
	// JwtIssuer         interfaces.JWTIssuer
	// JwtValidator      interfaces.JWTValidator
	Rbac *casbin.Enforcer
}

func (app *App) CommonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")

		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")

		next.ServeHTTP(w, r)
	})
}

func (app *App) RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a deferred function (which will always be run in the event
		// of a panic as Go unwinds the stack).
		defer func() {
			// Use the builtin recover function to check if there has been a
			// panic or not. If there has...
			if err := recover(); err != nil {
				// Set a "Connection: close" header on the response.
				w.Header().Set("Connection", "close")
				app.Logger.Error("recovered panic", slog.String("error", fmt.Sprintf("%v", err)))
				http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *App) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// var (
		// 	ip     = r.RemoteAddr
		// 	proto  = r.Proto
		// 	method = r.Method
		// 	uri    = r.URL.RequestURI()
		// )

		// app .App.Logger.Info("received request", "ip", ip, "proto", proto, "method", method, "uri", uri)

		next.ServeHTTP(w, r)
	})
}
