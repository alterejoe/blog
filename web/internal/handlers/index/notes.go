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
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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
		tools.Logger.Info("FleetingContainer", slog.String("fleetingTime", prefs.FleetingTime))
		props := htmxindex.FleetingProps{
			AlignInputTop: prefs.FleetingAlign == "Top",
			TimeFormat:    prefs.FleetingTime,
		}
		helpers.RenderContent(w, r, htmxindex.Fleeting(&props), htmxhtml.FullPage)
	}
}

type FleetingNoteRequest struct {
	NoteID string `validate:"required,uuid4"`
}

func FleetingNote(tools *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		params, err := validate.ParseAndValidate[FleetingNoteRequest](r)
		if err != nil {
			tools.Logger.Error("Error response", slog.Any("error", err.Error()))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var fleetingNote db.BlogFleeting

		err = tools.WithTx(r.Context(), app.RLSUser, func(q *db.Queries) error {
			nid, err := uuid.Parse(params.NoteID)
			if err != nil {
				return err
			}
			fleetingNote, err = q.GetFleetingNote(r.Context(), pgtype.UUID{
				Bytes: nid,
				Valid: true,
			})
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			tools.Logger.Error("Error response", slog.Any("error", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tools.Logger.Info("FleetingNote", slog.String("id", fleetingNote.ID.String()))

		props := htmxindex.FleetingNoteProps{
			Note: &fleetingNote,
		}

		htmxindex.FleetingNote(&props).Render(r.Context(), w)
	}
}

type FleetingNotesViewRequest struct {
}

func FleetingNotesViewContainer(tools *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_ = tools.Session.GetPreferences(r.Context())

		prefs := tools.Session.GetPreferences(r.Context())

		user_id := tools.Session.GetAuthUserID(r.Context())
		var fleetingNotes []db.BlogFleeting
		var err error

		err = tools.WithTx(r.Context(), app.RLSUser, func(q *db.Queries) error {
			fleetingNotes, err = q.GetFleetingNotesDesc(r.Context(), user_id)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			tools.Logger.Error("Error response", slog.Any("error", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tools.Logger.Info("FleetingNotesViewContainer", slog.Int("count", len(fleetingNotes)))
		props := htmxindex.FleetingNotesViewProps{
			Notes:         fleetingNotes,
			AlignInputTop: prefs.FleetingAlign == "Top",
		}

		helpers.RenderContent(w, r, htmxindex.FleetingNotesView(&props), htmxhtml.FullPage)
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
		var note db.BlogFleeting

		err = tools.WithTx(r.Context(), app.RLSUser, func(q *db.Queries) error {
			note, err = q.CreateFleetingNote(r.Context(), db.CreateFleetingNoteParams{
				UserID:   user_id,
				TagNames: []string{},
				Content:  params.Text,
			})
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			tools.Logger.Error("Error response", slog.Any("error", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// w.Header().Set("HX-Trigger", "fleeting-updated")
		// w.WriteHeader(http.StatusNoContent)
		htmxindex.FleetingNote(&htmxindex.FleetingNoteProps{
			Note: &note,
		}).Render(r.Context(), w)
	}
}

// --- Edit note content ---

type EditNoteRequest struct {
	Id      string `json:"id" validate:"required"`
	Content string `json:"content" validate:"required,min=1,max=10000"`
}

func EditNote(tools *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		noteID := chi.URLParam(r, "id")
		nid, err := uuid.Parse(noteID)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		params, err := validate.ParseAndValidate[EditNoteRequest](r)
		if err != nil {
			tools.Logger.Error("EditNote validation", slog.Any("error", err.Error()))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = tools.WithTx(r.Context(), app.RLSUser, func(q *db.Queries) error {
			return q.UpdateFleetingNoteContent(r.Context(), db.UpdateFleetingNoteContentParams{
				ID:      pgtype.UUID{Bytes: nid, Valid: true},
				Content: params.Content,
			})
		})
		if err != nil {
			tools.Logger.Error("EditNote", slog.Any("error", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// --- Soft delete ---

func SoftDeleteNote(tools *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		noteID := chi.URLParam(r, "id")
		nid, err := uuid.Parse(noteID)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		err = tools.WithTx(r.Context(), app.RLSUser, func(q *db.Queries) error {
			return q.SoftDeleteFleetingNote(r.Context(), pgtype.UUID{Bytes: nid, Valid: true})
		})
		if err != nil {
			tools.Logger.Error("SoftDeleteNote", slog.Any("error", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// --- Restore soft-deleted note ---

func RestoreNote(tools *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		noteID := chi.URLParam(r, "id")
		nid, err := uuid.Parse(noteID)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		err = tools.WithTx(r.Context(), app.RLSUser, func(q *db.Queries) error {
			return q.RestoreFleetingNote(r.Context(), pgtype.UUID{Bytes: nid, Valid: true})
		})
		if err != nil {
			tools.Logger.Error("RestoreNote", slog.Any("error", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
