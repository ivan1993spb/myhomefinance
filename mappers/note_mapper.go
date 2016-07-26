package mappers

import (
	"time"

	"github.com/ivan1993spb/myhomefinance/models"
)

type NoteMapper interface {
	CreateNote(time time.Time, name string, text string) (*models.Note, error)
	UpdateNote(id int64, time time.Time, name, text string) error
	DeleteNote(id int64) error
	GetNoteById(id int64) (*models.Note, error)
	GetNotesByTimeRange(from time.Time, to time.Time) ([]*models.Note, error)
	GetNotesByTimeRangeGrep(from time.Time, to time.Time, name string) ([]*models.Note, error)
}

type ErrCreateNoteMapper string

func (e ErrCreateNoteMapper) Error() string {
	return "cannot creating note mapper: " + string(e)
}

type ErrCreateNote string

func (e ErrCreateNote) Error() string {
	return "cannot create note: " + string(e)
}

type ErrDeleteNote string

func (e ErrDeleteNote) Error() string {
	return "cannot delete note: " + string(e)
}

var ErrDeleteNoteNotFound = ErrDeleteNote("note to delete was not found")

type ErrUpdateNote string

func (e ErrUpdateNote) Error() string {
	return "cannot update note: " + string(e)
}

var ErrUpdateNoteNotFound = ErrUpdateNote("note to update was not found")

type ErrGetNoteById string

func (e ErrGetNoteById) Error() string {
	return "cannot get note by id: " + string(e)
}

var ErrGetNoteByIdNotFound = ErrGetNoteById("note with given id was not found")

type ErrGetNotesByTimeRange string

func (e ErrGetNotesByTimeRange) Error() string {
	return "cannot get notes by time range: " + string(e)
}

var ErrGetNotesByTimeRangeInvalidTimeRange = ErrGetNotesByTimeRange("invalid time range")

type ErrGetNotesByTimeRangeMatch string

func (e ErrGetNotesByTimeRangeMatch) Error() string {
	return "cannot get notes by time range and match: " + string(e)
}

var (
	ErrGetNotesByTimeRangeMatchEmptyName        = ErrGetNotesByTimeRangeMatch("passed empty note name for matching")
	ErrGetNotesByTimeRangeMatchInvalidTimeRange = ErrGetNotesByTimeRangeMatch("invalid time range")
)
