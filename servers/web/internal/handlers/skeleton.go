package handlers

import (
	"blog/servers/web/internal/app"
	"blog/servers/web/ui/html"
	"net/http"
)

func Pages(s string) string {
	switch s {
	case "index":
		return "index"
	case "blog":
		return "blog"
	case "theme":
		return "theme"
	default:
		return "index"
	}
}

func Skeleton(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		page := r.URL.Query().Get("page")
		_, _ = app.Session.GetFlashMessage(r.Context())

		// page = ValidatePage(page)

		props := &html.SkeletonProps{
			// Page: "landing",
			Page: Pages(page),
		}

		html.Skeleton(props).Render(r.Context(), w)
	}
}
