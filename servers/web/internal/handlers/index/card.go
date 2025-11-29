package index

import (
	"blog/servers/web/internal/app"
	"blog/servers/web/ui/html/layouts"
	"log/slog"
	"net/http"
)

func Orientations(s string) string {
	switch s {
	case "left":
		return "left"
	case "right":
		return "right"
	default:
		return "left"
	}
}
func Card(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		page := r.URL.Query().Get("orientation")
		id := r.URL.Query().Get("id")
		_, _ = app.Session.GetFlashMessage(r.Context())

		// page = ValidatePage(page)

		app.Logger.Info("Card", slog.Any("page", page))
		props := &layouts.CardProps{
			Id:          id,
			Orientation: Orientations(page),
		}
		app.Logger.Info("Card", slog.Any("props", props))
		layouts.CardOpen(props).Render(r.Context(), w)
	}
}
