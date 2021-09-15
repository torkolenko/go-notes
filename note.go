package apiServer

import "errors"

type Note struct {
	Id    int    `json:"id"`
	Title string `json:"title" db:"title" binding:"required"`
	Text  string `json:"text" db:"text"`
}

type UpdateNoteInput struct {
	Title *string `json:"title"`
	Text  *string `json:"text"`
}

func (i UpdateNoteInput) Validate() error {
	if i.Title == nil && i.Text == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
