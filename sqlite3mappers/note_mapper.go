package sqlite3mappers

import (
	"database/sql"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/ivan1993spb/myhomefinance/models"
)

type ErrNoteMapper string

func (e ErrNoteMapper) Error() string {
	return "note mapper error: " + string(e)
}

type NoteMapper struct {
	*sql.DB
}

func (nm *NoteMapper) CreateNote(t time.Time, name, text string) (*models.Note, error) {
	if len(name) == 0 {
		return nil, ErrNoteMapper("name cannot be empty")
	}

	res, err := nm.DB.Exec("INSERT INTO `notes` (`name`, `unixtimestamp`, `text`) VALUES (?, ?, ?)",
		name, t.Unix(), text)
	if err != nil {
		return nil, ErrNoteMapper("cannot insert new note into db: " + err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, ErrOutflowMapper("cannot get id of new note: " + err.Error())
	}

	return &models.Note{
		ID:       id,
		Datetime: &strfmt.DateTime(t),
		Name:     &name,
		Text:     text,
	}, nil
}

func (nm *NoteMapper) DeleteNote(id uint64) error {
	result, err := nm.DB.Exec("DELETE FROM `notes` WHERE `id` = ?", id)
	if err != nil {
		return ErrNoteMapper("cannot delete note: " + err.Error())
	}

	n, _ := result.RowsAffected()
	if n == 0 {
		return ErrNoteMapper("cannot delete note: note not found")
	}

	return nil
}

func (nm *NoteMapper) UpdateNote(note *models.Note) error {
	result, err := nm.DB.Exec("UPDATE `notes` SET `name` = ?, `unixtimestamp` = ?, `text` = ? WHERE `id` = ?",
		note.Name, note.Datetime.String(), note.Text, note.ID)
	if err != nil {
		return ErrNoteMapper("cannot update note: " + err.Error())
	}

	n, _ := result.RowsAffected()
	if n == 0 {
		return ErrNoteMapper("cannot update note: note not found")
	}

	return nil
}

func (nm *NoteMapper) GetNoteById(id uint64) (*models.Note, error) {
	var (
		note          = &models.Note{}
		unixtimestamp int64
	)

	err := nm.DB.QueryRow("SELECT `id`, `unixtimestamp`, `name`, `text` FROM `notes` WHERE `id` = ?", id).
		Scan(&note.ID, &unixtimestamp, note.Name, &note.Text)
	if err != nil {
		return nil, ErrNoteMapper("cannot get note by id: " + err.Error())
	}

	note.Datetime = &strfmt.DateTime(time.Unix(unixtimestamp, 0))

	return note, nil
}

func (nm *NoteMapper) GetNotesByTimeRange(from time.Time, to time.Time) ([]*models.Note, error) {
	if from.Unix() >= to.Unix() {
		return nil, ErrNoteMapper("invalid time range")
	}

	rows, err := nm.DB.Query("SELECT `id`, `unixtimestamp`, `name`, `text` FROM `notes` "+
		"WHERE `unixtimestamp` BETWEEN ? AND ?", from.Unix(), to.Unix())
	if err != nil {
		return nil, ErrNoteMapper("cannot get notes by time range: " + err.Error())
	}
	defer rows.Close()

	var notes = make([]*models.Note, 0)

	for rows.Next() {
		var (
			note          = &models.Note{}
			unixtimestamp int64
		)
		if err := rows.Scan(&note.ID, &unixtimestamp, note.Name, &note.Text); err != nil {
			return nil, ErrNoteMapper("cannot get notes by time range: " + err.Error())
		}
		note.Datetime = &strfmt.DateTime(time.Unix(unixtimestamp, 0))
		notes = append(notes, note)
	}
	if err := rows.Err(); err != nil {
		return nil, ErrNoteMapper("cannot get notes by time range: " + err.Error())
	}

	return notes, nil
}
