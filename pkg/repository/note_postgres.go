package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	apiServer "github.com/torkolenko/go-notes"
)

type NotePostgres struct {
	db *sqlx.DB
}

func NewNotePostgres(db *sqlx.DB) *NotePostgres {
	return &NotePostgres{db: db}
}

func (r *NotePostgres) CreateNote(note apiServer.Note) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, text) VALUES ($1, $2) RETURNING id", notesTable)

	row := r.db.QueryRow(query, note.Title, note.Text)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *NotePostgres) GetAll() ([]apiServer.Note, error) {
	var notes []apiServer.Note

	query := fmt.Sprintf("SELECT * FROM %s", notesTable)

	err := r.db.Select(&notes, query)

	return notes, err
}

func (r *NotePostgres) GetById(noteId int) (apiServer.Note, error) {
	var note apiServer.Note

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", notesTable)

	err := r.db.Get(&note, query, noteId)

	return note, err
}

func (r *NotePostgres) Delete(noteId int) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", notesTable)

	_, err := r.db.Exec(query, noteId)

	return err
}

func (r *NotePostgres) Update(noteId int, input apiServer.UpdateNoteInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Text != nil {
		setValues = append(setValues, fmt.Sprintf("text=$%d", argId))
		args = append(args, *input.Text)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", notesTable, setQuery, argId)
	args = append(args, noteId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("updateQuery: %s", args)

	_, err := r.db.Exec(query, args...)

	return err
}
