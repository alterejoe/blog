package main

import (
	"net/http"
	"os"

	"github.com/alterejoe/blog/internal/app"
	"github.com/alterejoe/blog/internal/handlers/auth"
	"github.com/alterejoe/blog/internal/handlers/index"

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

	serveStaticFile("/static/css/output.css", "./ui/static/css/output.css", r)
	serveStaticFile("/static/js/htmx.min.js", "./ui/static/js/htmx.min.js", r)
	serveStaticFile("/static/js/landing.js", "./ui/static/js/landing.js", r)

	serveStaticFile("/static/js/form-behavior/landing.js", "./ui/static/js/form-behavior/landing.js", r)
	serveStaticFile("/static/js/form-behavior/constraints.js", "/static/js/form-behavior/constraints.js", r)
	serveStaticFile("/static/js/form-behavior/core.js", "/static/js/form-behavior/core.js", r)
	serveStaticFile("/static/js/form-behavior/dirty-tracking.js", "/static/js/form-behavior/dirty-tracking.js", r)
	serveStaticFile("/static/js/form-behavior/init.js", "/static/js/form-behavior/init.js", r)
	serveStaticFile("/static/js/form-behavior/sanitizers.js", "/static/js/form-behavior/sanitizers.js", r)

	r.Route("/auth", func(r chi.Router) {
		// r.Get("/landing", auth.Landing(app))
		r.Get("/login", auth.Login(app))
		r.Get("/logout", auth.Logout(app))
		r.Get("/callback", auth.Callback(app))
	})

	r.Group(func(r chi.Router) {
		r.Use(ClientAuthMiddleware(app))
		r.Get("/", http.RedirectHandler("/index", http.StatusSeeOther).ServeHTTP)
		r.Get("/index", index.Index(app))
		r.Get("/blog", index.Blog(app))
		r.Get("/notes", index.Notes(app))
		r.Get("/theme", index.Theme(app))

		r.Get("/nav", index.Nav(app))
		if os.Getenv("ENVIRONMENT") == "dev" {
			r.Get("/debug", index.Debug(app))
			r.Get("/debug/empty", index.DebugEmpty(app))
		}
	})
	r.Group(func(r chi.Router) {
		r.Route("/admin", func(r chi.Router) {
			r.Get("/", index.Blog(app))
		})
	})

	return r
}
