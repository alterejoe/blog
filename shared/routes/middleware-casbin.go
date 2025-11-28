package routes

import (
	"net/http"

	"github.com/casbin/casbin/v2"
)

func CasbinMiddleware(e *casbin.Enforcer, getSub func(r *http.Request) string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sub := getSub(r)
			obj := r.URL.Path
			act := r.Method

			ok, err := e.Enforce(sub, obj, act)
			if err != nil {
				http.Error(w, "authorization error", http.StatusInternalServerError)
				return
			}

			if !ok {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
