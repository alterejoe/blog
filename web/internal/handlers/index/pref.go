package index

import (
	"log/slog"
	"net/http"
	"slices"

	"github.com/alterejoe/blog/internal/app"
	"github.com/go-chi/chi/v5"
)

type Preference struct {
	Values []string
	Event  string
}

var preferences = map[string]Preference{
	"fleeting_align": {
		Values: []string{"Top", "Bottom"},
		Event:  "fleeting-updated",
	},
	"fleeting_time": {
		Values: []string{"Absolute", "Compact", "Verbose"},
		Event:  "fleeting-updated",
	},
}

func SetPreference(tools *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "key")
		value := chi.URLParam(r, "value")

		pref, ok := preferences[key]
		if !ok {
			http.Error(w, "invalid preference key", http.StatusBadRequest)
			return
		}

		if !slices.Contains(pref.Values, value) {
			http.Error(w, "invalid preference value", http.StatusBadRequest)
			return
		}

		tools.Session.SetPreference(r.Context(), key, value)
		tools.Logger.Info("SetPreference", slog.String("key", key), slog.String("value", value))

		if pref.Event != "" {
			w.Header().Set("HX-Trigger", pref.Event)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
