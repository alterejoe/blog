package index

import (
	"github.com/alterejoe/blog/internal/app"
	"net/http"
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
