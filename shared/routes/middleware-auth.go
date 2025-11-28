package routes

import (
	"net/http"
)

func (app *Dependencies[T]) IsAuthenticated(r *http.Request) bool {
	user := app.Session().Get(r.Context(), "authenticatedUserID")
	uid, ok := user.(string)
	if !ok || uid == "" {
		// app.Logger().Info("unauthenticated or empty user id")
		return false
	}
	// app.Logger().Debug("User is authenticated")
	return true
}

func (app *Dependencies[T]) RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If the user is not authenticated, redirect them to the login page and
		// return from the middleware chain so that no subsequent handlers in
		// the chain are executed.

		// need to check to see if user is still associated with group_id in session

		if !app.IsAuthenticated(r) {
			app.Logger().Debug("User is not authenticated")
			app.Session().Destroy(r.Context())
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}
		// is isLoggedIn

		// Otherwise set the "Cache-Control: no-store" header so that pages
		// require authentication are not stored in the users browser cache (or
		// other intermediary cache).
		w.Header().Add("Cache-Control", "no-store")

		// And call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}
