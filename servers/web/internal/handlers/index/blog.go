package index

import (
	"blog/servers/web/internal/app"
	htmxindex "blog/servers/web/ui/html/index"
	"net/http"
)

func Blog(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		// page := r.URL.Query().Get("page")
		_, _ = app.Session.GetFlashMessage(r.Context())

		// page = ValidatePage(page)

		props := &htmxindex.IndexProps{}

		htmxindex.Index(props).Render(r.Context(), w)
	}
}
