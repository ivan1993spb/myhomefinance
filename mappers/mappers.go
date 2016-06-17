package mappers

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ivan1993spb/myhomefinance/models"
)

var ErrFindNoteById = errors.New("cannot find note with passed id")

type InitDB func() (*sql.DB, error)

type FinalizeDB func() error

type InflowMapper interface {
	CreateInflow(t time.Time, name string, amount float64, description, source string) (*models.Inflow, error)
}

type OutflowMapper interface {
	CreateOutflow(t time.Time, name string, amount float64, description, destination, target string, count float64,
		metricUnit string, satisfaction float32) (*models.Outflow, error)
}

type NoteMapper interface {
	CreateNote(time time.Time, name string, text string) (*models.Note, error)
	UpdateNote(id int64, time time.Time, name, text string) error
	DeleteNote(id int64) error
	GetNoteById(id int64) (*models.Note, error)
	GetNotesByTimeRange(from time.Time, to time.Time) ([]*models.Note, error)
	GetNotesByTimeRangeGrep(from time.Time, to time.Time, name string) ([]*models.Note, error)
}

type HistoryMapper interface {
}
