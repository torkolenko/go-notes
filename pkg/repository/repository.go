package repository

import (
	"github.com/jmoiron/sqlx"
	apiServer "github.com/torkolenko/go-notes"
)

type Note interface {
	CreateNote(note apiServer.Note) (int, error)
	GetAll() ([]apiServer.Note, error)
	GetById(noteId int) (apiServer.Note, error)
	Delete(noteId int) error
	Update(noteId int, input apiServer.UpdateNoteInput) error
}

type Repository struct {
	Note
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Note: NewNotePostgres(db),
	}
}
