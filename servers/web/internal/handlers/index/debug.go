package index

import (
	"net/http"
	"zotes/servers/web/internal/app"
)

// func Skeleton(w http.ResponseWriter, r *http.Request, app interfaces.BasicDep) {

func Debug(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Debug"))
	}
}

func DebugEmpty(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	}
}
