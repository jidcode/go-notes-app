package routes

import (
	"net/http"

	"github.com/code/go-project/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(noteHandler *handlers.NoteHandler) *chi.Mux {
	r := chi.NewRouter()

	// Adding middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// Define the endpoint routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Notes App!!"))
	})

	r.Route("/api/notes", func(r chi.Router) {
		r.Post("/", noteHandler.CreateNote)
		r.Get("/", noteHandler.GetNotes)
		r.Get("/{id}", noteHandler.GetNoteByID)
		r.Put("/{id}", noteHandler.UpdateNote)
		r.Delete("/{id}", noteHandler.DeleteNote)
	})

	return r
}
