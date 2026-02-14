package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/alterejoe/blog/internal/app"
	"github.com/alterejoe/blog/internal/handlers/auth"
	"github.com/alterejoe/blog/internal/handlers/index"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func serveStaticFile(routePath, diskPath string, r chi.Router) {
	r.Get(routePath, func(w http.ResponseWriter, req *http.Request) {
		// 1. Check file exists
		info, err := os.Stat(diskPath)
		if err != nil {
			if os.IsNotExist(err) {
				log.Printf("[static] NOT FOUND: %s -> %s", routePath, diskPath)
				http.NotFound(w, req)
				return
			}
			log.Printf("[static] ERROR stating file %s: %v", diskPath, err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		// 2. Prevent directory serving
		if info.IsDir() {
			log.Printf("[static] REFUSED directory: %s", diskPath)
			http.Error(w, "not a file", http.StatusBadRequest)
			return
		}

		// 3. Force correct content-type for SVG (critical for favicons)
		if strings.HasSuffix(diskPath, ".svg") {
			w.Header().Set("Content-Type", "image/svg+xml")
		}

		// 4. Optional: disable aggressive caching while debugging
		w.Header().Set("Cache-Control", "no-store")

		log.Printf("[static] SERVING %s (%d bytes)", diskPath, info.Size())

		http.ServeFile(w, req, diskPath)
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
	// serveStaticFile("/static/js/landing.js", "./ui/static/js/landing.js", r)
	serveStaticFile("/static/js/time.js", "./ui/static/js/time.js", r)

	serveStaticFile("/static/js/form-behavior/constraints.js", "./ui/static/js/form-behavior/constraints.js", r)
	serveStaticFile("/static/js/form-behavior/core.js", "./ui/static/js/form-behavior/core.js", r)
	serveStaticFile("/static/js/form-behavior/dirty-tracking.js", "./ui/static/js/form-behavior/dirty-tracking.js", r)
	serveStaticFile("/static/js/form-behavior/init.js", "./ui/static/js/form-behavior/init.js", r)
	serveStaticFile("/static/js/form-behavior/sanitizers.js", "./ui/static/js/form-behavior/sanitizers.js", r)

	serveStaticFile("/static/js/vim/state.js", "./ui/static/js/vim/state.js", r)
	serveStaticFile("/static/js/vim/keymap.js", "./ui/static/js/vim/keymap.js", r)
	serveStaticFile("/static/js/vim/dom.js", "./ui/static/js/vim/dom.js", r)
	serveStaticFile("/static/js/vim/modes.js", "./ui/static/js/vim/modes.js", r)
	serveStaticFile("/static/js/vim/movement.js", "./ui/static/js/vim/movement.js", r)
	serveStaticFile("/static/js/vim/actions.js", "./ui/static/js/vim/actions.js", r)
	serveStaticFile("/static/js/vim/keys.js", "./ui/static/js/vim/keys.js", r)
	serveStaticFile("/static/js/vim/init.js", "./ui/static/js/vim/init.js", r)

	serveStaticFile("/static/favicon.svg", "./ui/static/favicon1.svg", r)

	r.Route("/auth", func(r chi.Router) {
		// r.Get("/landing", auth.Landing(app))
		r.Get("/login", auth.Login(app))
		r.Get("/logout", auth.Logout(app))
		r.Get("/callback", auth.Callback(app))
	})

	r.Group(func(r chi.Router) {
		r.Use(ClientAuthMiddleware(app))
		r.Route("/nav", func(r chi.Router) {
			r.Get("/", index.Nav(app))
			r.Get("/logout", index.Logout(app))
		})
		r.Get("/", index.Index(app))
		r.Get("/blog", index.Blog(app))
		r.Get("/theme", index.Theme(app))
		r.Route("/notes", func(r chi.Router) {
			r.Get("/", index.Notes(app))
			r.Route("/fleeting", func(r chi.Router) {
				r.Get("/", index.FleetingContainer(app))
				r.Get("/view", index.FleetingNotesViewContainer(app))
				r.Post("/new", index.PostNewNote(app))
				r.Route("/{id}", func(r chi.Router) {
					r.Get("/edit", index.EditNote(app))
					r.Patch("/edit", index.EditNote(app))
					r.Patch("/delete", index.SoftDeleteNote(app))
					r.Patch("/restore", index.RestoreNote(app))

				})
			})
		})
		if os.Getenv("ENVIRONMENT") == "dev" {
			r.Get("/debug", index.Debug(app))
			r.Get("/debug/empty", index.DebugEmpty(app))
		}
		r.Put("/pref/{key}/{value}", index.SetPreference(app))

		r.Group(func(r chi.Router) {
			r.Route("/admin", func(r chi.Router) {
				r.Get("/", index.Blog(app))
			})
		})
	})

	return r
}
