package service

import (
	apiServer "github.com/torkolenko/go-notes"
	"github.com/torkolenko/go-notes/pkg/repository"
)

type Note interface {
	CreateNote(note apiServer.Note) (int, error)
	GetAll() ([]apiServer.Note, error)
	GetById(noteId int) (apiServer.Note, error)
	Delete(noteId int) error
	Update(noteId int, input apiServer.UpdateNoteInput) error
}

type Service struct {
	Note
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Note: NewNoteService(repos.Note),
	}
}
