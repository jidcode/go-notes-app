package repository

import (
	"github.com/code/go-project/models"
	"github.com/jmoiron/sqlx"
)

// NoteRepository handles the CRUD operations for notes
type NoteRepository struct {
	db *sqlx.DB
}

// NewNoteRepository returns a new instance of NoteRepository
func NewNoteRepository(db *sqlx.DB) *NoteRepository {
	return &NoteRepository{db: db}
}

// CreateNote adds a new note to the database
func (r *NoteRepository) CreateNote(note *models.Note) error {
	query := `INSERT INTO notes (title, content, created_at, updated_at) 
              VALUES ($1, $2, NOW(), NOW()) RETURNING id, created_at, updated_at`

	return r.db.QueryRowx(query, note.Title, note.Content).Scan(&note.ID, &note.CreatedAt, &note.UpdatedAt)
}

// GetNotes retrieves all notes from the database
func (r *NoteRepository) GetNotes() ([]models.Note, error) {
	notes := []models.Note{}
	query := `SELECT * FROM notes ORDER BY created_at DESC`
	err := r.db.Select(&notes, query)
	return notes, err
}

// GetNoteByID retrieves a note by its ID
func (r *NoteRepository) GetNoteByID(id int) (*models.Note, error) {
	note := &models.Note{}
	query := `SELECT * FROM notes WHERE id = $1`
	err := r.db.Get(note, query, id)
	return note, err
}

// UpdateNote modifies the content of an existing note
func (r *NoteRepository) UpdateNote(note *models.Note) error {
	query := `UPDATE notes SET title = $1, content = $2, updated_at = NOW() WHERE id = $3`
	_, err := r.db.Exec(query, note.Title, note.Content, note.ID)
	return err
}

// DeleteNote removes a note from the database
func (r *NoteRepository) DeleteNote(id int) error {
	query := `DELETE FROM notes WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
