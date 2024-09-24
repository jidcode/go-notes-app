package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/code/go-project/models"
	"github.com/code/go-project/repository"
	"github.com/go-chi/chi/v5"
)

// NoteHandler handles HTTP requests related to notes
type NoteHandler struct {
	Repo *repository.NoteRepository
}

// NewNoteHandler returns a new NoteHandler
func NewNoteHandler(repo *repository.NoteRepository) *NoteHandler {
	return &NoteHandler{Repo: repo}
}

// CreateNote creates a new note
func (h *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.Repo.CreateNote(&note); err != nil {
		http.Error(w, "Error creating note", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// GetNotes retrieves all notes
func (h *NoteHandler) GetNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := h.Repo.GetNotes()
	if err != nil {
		http.Error(w, "Error fetching notes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

// GetNoteByID retrieves a note by its ID
func (h *NoteHandler) GetNoteByID(w http.ResponseWriter, r *http.Request) {
	noteId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(noteId)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	note, err := h.Repo.GetNoteByID(id)
	if err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// UpdateNote updates an existing note
func (h *NoteHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	noteId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(noteId)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	note.ID = id
	if err := h.Repo.UpdateNote(&note); err != nil {
		http.Error(w, "Error updating note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteNote deletes a note by its ID
func (h *NoteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	noteId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(noteId)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	if err := h.Repo.DeleteNote(id); err != nil {
		http.Error(w, "Error deleting note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
