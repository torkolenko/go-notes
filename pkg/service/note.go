package service

import (
	apiServer "github.com/torkolenko/go-notes"
	"github.com/torkolenko/go-notes/pkg/repository"
)

type NoteService struct {
	repo repository.Note
}

func NewNoteService(repo repository.Note) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) CreateNote(note apiServer.Note) (int, error) {
	return s.repo.CreateNote(note)
}

func (s *NoteService) GetAll() ([]apiServer.Note, error) {
	return s.repo.GetAll()
}

func (s *NoteService) GetById(noteId int) (apiServer.Note, error) {
	return s.repo.GetById(noteId)
}

func (s *NoteService) Delete(noteId int) error {
	return s.repo.Delete(noteId)
}

func (s *NoteService) Update(noteId int, input apiServer.UpdateNoteInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(noteId, input)
}
