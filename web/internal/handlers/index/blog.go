package index

import (
	"net/http"

	"github.com/alterejoe/blog/internal/app"
	htmxhtml "github.com/alterejoe/blog/ui/html"
	htmxindex "github.com/alterejoe/blog/ui/html/index"
)

func Blog(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		// page := r.URL.Query().Get("page")
		_, _ = app.Session.GetFlashMessage(r.Context())

		props := &htmxindex.IndexProps{}

		content := htmxindex.Blog(props)
		htmxhtml.FullPage(content).Render(r.Context(), w)
	}
}
