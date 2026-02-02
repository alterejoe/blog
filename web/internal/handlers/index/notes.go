package index

import (
	"log/slog"
	"net/http"

	"github.com/alterejoe/blog/db"
	"github.com/alterejoe/blog/internal/app"
	"github.com/alterejoe/blog/internal/helpers"
	htmxhtml "github.com/alterejoe/blog/ui/html"
	htmxindex "github.com/alterejoe/blog/ui/html/index"
	"github.com/alterejoe/shared/validate"
)

type NotesRequest struct {
}

func Notes(tools *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		props := htmxindex.NotesProps{
			FolderStructure: &htmxindex.FolderStructureProps{},
		}

		// content := htmxindex.Blog(props)
		// htmxhtml.FullPage(content).Render(r.Context(), w)
		helpers.RenderContent(w, r, htmxindex.Notes(&props), htmxhtml.FullPage)
	}
}

type FleetingContainerRequest struct {
}

func FleetingContainer(tools *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		prefs := tools.Session.GetPreferences(r.Context())

		props := htmxindex.FleetingProps{
			AlignInputTop: prefs.FleetingAlign == "Top",
		}

		helpers.RenderContent(w, r, htmxindex.Fleeting(&props), htmxhtml.FullPage)
	}
}

type PostNewNoteRequest struct {
	Text string `form:"fleeting-post" validate:"required,min=1,max=10000"`
}

func PostNewNote(tools *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		params, err := validate.ParseAndValidate[PostNewNoteRequest](r)
		if err != nil {
			tools.Logger.Error("Error response", slog.Any("error", err.Error()))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user_id := tools.Session.GetAuthUserID(r.Context())
		tools.Logger.Info("PostNewNote", slog.String("user_id", user_id.String()), slog.String("text", params.Text))

		err = tools.WithTx(r.Context(), app.RLSUser, func(q *db.Queries) error {
			_, err := q.CreateFleetingNote(r.Context(), db.CreateFleetingNoteParams{
				UserID:  user_id,
				TagName: params.Text,
			})
			if err != nil {
				return err
			}
			return nil
		})

		// helpers.RenderContent(w, r, htmxindex.PostNewNote(), htmxhtml.FullPage)
	}
}
