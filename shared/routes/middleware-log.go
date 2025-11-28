package routes

import (
	"net/http"
)

func (app *Dependencies[T]) LogRequest(next http.Handler) http.Handler {
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

// TODO: OPS-10
