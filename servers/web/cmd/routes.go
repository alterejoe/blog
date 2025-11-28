package main

import (
	"net/http"
	"os"

	"zotes/servers/web/internal/app"
	"zotes/servers/web/internal/handlers/auth"
	"zotes/servers/web/internal/handlers/index"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func serveStaticFile(path, disk string, r chi.Router) {
	r.Get(path, func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, disk)
	})
}

func GetRoutes(app *app.App) *chi.Mux {
	r := chi.NewRouter()
	// fileserver := http.FileServer(routes.NeuteredFileSystem{Fs: http.Dir("./ui/static")})
	//
	r.Use(app.Session.LoadAndSave, app.RecoverPanic, app.LogRequest, app.CommonHeaders)
	r.Use(middleware.RedirectSlashes)

	serveStaticFile("/static/js/file-upload.js", "./ui/static/js/file-upload.js", r)
	serveStaticFile("/static/js/file-staging.js", "./ui/static/js/file-staging.js", r)
	serveStaticFile("/static/css/output.css", "./ui/static/css/output.css", r)
	serveStaticFile("/static/js/htmx.min.js", "./ui/static/js/htmx.min.js", r)
	serveStaticFile("/static/js/checkboxes.js", "./ui/static/js/checkboxes.js", r)
	serveStaticFile("/static/js/state.js", "./ui/static/js/state.js", r)

	r.Route("/auth", func(r chi.Router) {
		r.Get("/landing", auth.Landing(app))
		r.Get("/login", auth.Login(app))
		r.Get("/logout", auth.Logout(app))
		r.Get("/callback", auth.Callback(app))
	})

	r.Group(func(r chi.Router) {
		r.Use(ClientAuthMiddleware(app))
		r.Get("/", index.Skeleton(app))
		if os.Getenv("ENVIRONMENT") == "dev" {
			r.Get("/debug", index.Debug(app))
			r.Get("/debug/empty", index.DebugEmpty(app))
		}
	})

	return r
}
