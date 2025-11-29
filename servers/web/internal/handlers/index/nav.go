package index

import (
	"blog/servers/web/internal/app"
	layouts "blog/servers/web/ui/html/layouts"
	"net/http"
)

func Nav(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		// page := r.URL.Query().Get("page")
		_, _ = app.Session.GetFlashMessage(r.Context())

		// page = ValidatePage(page)

		props := &layouts.NavProps{}

		layouts.NavOpen(props).Render(r.Context(), w)
	}
}
