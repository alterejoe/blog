package index

import (
	"blog/servers/web/internal/app"
	htmxindex "blog/servers/web/ui/html/index"
	"encoding/gob"
	"log/slog"
	"net/http"
)

func init() {
	gob.Register(map[string]string{})
}

func getOrCreateTheme(app *app.App, r *http.Request) *map[string]string {
	// Get map from session instead of struct
	colorsMap := app.Session.Get(r.Context(), "user_theme_colors")

	var colors map[string]string
	if colorsMap != nil {
		if cm, ok := colorsMap.(map[string]string); ok {
			colors = cm
		}
	}

	if colors == nil {
		colors = map[string]string{
			"primary":   "#3b82f6",
			"secondary": "#8b5cf6",
		}
		app.Session.Put(r.Context(), "user_theme_colors", colors)
	}

	return &colors
}

func Theme(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		// Get or create theme from session
		theme := getOrCreateTheme(app, r)
		app.Logger.Info("Theme", slog.Any("theme", theme))

		props := &htmxindex.ThemeProps{
			UserTheme: theme,
		}
		htmxindex.Theme(props).Render(r.Context(), w)
	}
}
