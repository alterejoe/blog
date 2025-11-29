package handlers

import (
	"blog/servers/web/internal/app"
	"blog/servers/web/ui/html"
	"net/http"
)

func Skeleton(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		// page := r.URL.Query().Get("page")
		_, _ = app.Session.GetFlashMessage(r.Context())

		// page = ValidatePage(page)

		props := &html.SkeletonProps{
			// Page: "landing",
			Page: "index",
		}

		html.Skeleton(props).Render(r.Context(), w)
	}
}
