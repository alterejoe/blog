package index

import (
	"net/http"

	"github.com/alterejoe/blog/internal/app"
	layouts "github.com/alterejoe/blog/ui/html/layouts"
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
