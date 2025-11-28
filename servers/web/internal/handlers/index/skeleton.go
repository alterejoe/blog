package index

import (
	"net/http"
	"zotes/servers/web/internal/app"
	"zotes/servers/web/ui/html"
)

func Skeleton(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		// page := r.URL.Query().Get("page")
		_, _ = app.Session.GetFlashMessage(r.Context())

		// page = ValidatePage(page)

		props := &html.SkeletonProps{}

		html.Skeleton(props).Render(r.Context(), w)
	}
}
