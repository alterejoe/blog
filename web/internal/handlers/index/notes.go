package index

import (
	"net/http"

	"github.com/alterejoe/blog/internal/app"
	"github.com/alterejoe/blog/internal/helpers"
	htmxhtml "github.com/alterejoe/blog/ui/html"
	htmxindex "github.com/alterejoe/blog/ui/html/index"
)

func Notes(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		props := htmxindex.NotesProps{
			Fleeting:        &htmxindex.FleetingProps{},
			FolderStructure: &htmxindex.FolderStructureProps{},
		}

		// content := htmxindex.Blog(props)
		// htmxhtml.FullPage(content).Render(r.Context(), w)
		helpers.RenderContent(w, r, htmxindex.Notes(props), htmxhtml.FullPage)
	}
}
