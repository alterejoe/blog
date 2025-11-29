package auth

import (
	"net/http"

	"blog/servers/web/internal/app"
	"blog/servers/web/ui/html"
)

func Landing(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		props := html.LandingProps{}
		html.Landing(&props).Render(r.Context(), w)
	}
}
