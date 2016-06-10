package sqlite3mappers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/ivan1993spb/myhomefinance/mappers"
	"github.com/ivan1993spb/myhomefinance/models"
)

const (
	_SQL_INSERT_NOTE       = "INSERT INTO `notes` (`name`, `unixtimestamp`, `text`) VALUES (?, ?, ?)"
	_SQL_DELETE_NOTE_BY_ID = "DELETE FROM `notes` WHERE `id` = ?"
	_SQL_SELECT_NOTE_BY_ID = "SELECT `id`, `unixtimestamp`, `name`, `text` FROM `notes` WHERE `id` = ?"
	_SQL_UPDATE_NOTE_BY_ID = "UPDATE `notes` SET `name` = ?, `unixtimestamp` = ?, `text` = ? WHERE `id` = ?"

	_SQL_SELECT_NOTES_BY_TIME_RANGE = "SELECT `id`, `unixtimestamp`, `name`, `text` FROM `notes` " +
		"WHERE `unixtimestamp` BETWEEN ? AND ?"
	_SQL_SELECT_NOTES_BY_TIME_RANGE_GREP = "SELECT `id`, `unixtimestamp`, `name`, `text` FROM `notes` " +
		"WHERE `unixtimestamp` BETWEEN ? AND ? AND grep(`name`, ?)"
)

type NoteMapper struct {
	db *sql.DB

	// Prepared statements

	insertNote *sql.Stmt
	deleteNote *sql.Stmt

	selectNoteById *sql.Stmt
	updateNoteById *sql.Stmt

	selectNotesByTimeRange     *sql.Stmt
	selectNotesByTimeRangeGrep *sql.Stmt
}

type errCreateNoteMapper string

func (e errCreateNoteMapper) Error() string {
	return "error creating note mapper: " + string(e)
}

func NewNoteMapper(db *sql.DB) (*NoteMapper, error) {
	err := db.Ping()
	if err != nil {
		return nil, errCreateNoteMapper(err.Error())
	}

	noteMapper := &NoteMapper{db: db}

	noteMapper.insertNote, err = db.Prepare(_SQL_INSERT_NOTE)
	if err != nil {
		return nil, errCreateNoteMapper(err.Error())
	}

	noteMapper.deleteNote, err = db.Prepare(_SQL_DELETE_NOTE_BY_ID)
	if err != nil {
		return nil, errCreateNoteMapper(err.Error())
	}

	noteMapper.selectNoteById, err = db.Prepare(_SQL_SELECT_NOTE_BY_ID)
	if err != nil {
		return nil, errCreateNoteMapper(err.Error())
	}

	noteMapper.updateNoteById, err = db.Prepare(_SQL_UPDATE_NOTE_BY_ID)
	if err != nil {
		return nil, errCreateNoteMapper(err.Error())
	}

	noteMapper.selectNotesByTimeRange, err = db.Prepare(_SQL_SELECT_NOTES_BY_TIME_RANGE)
	if err != nil {
		return nil, errCreateNoteMapper(err.Error())
	}

	noteMapper.selectNotesByTimeRangeGrep, err = db.Prepare(_SQL_SELECT_NOTES_BY_TIME_RANGE_GREP)
	if err != nil {
		return nil, errCreateNoteMapper(err.Error())
	}

	return noteMapper, nil
}

type errCreateNote string

func (e errCreateNote) Error() string {
	return "cannot create note: " + string(e)
}

func (nm *NoteMapper) CreateNote(datetime strfmt.DateTime, name string, text string) (*models.Note, error) {
	if len(name) == 0 {
		return nil, errCreateNote("name cannot be empty")
	}

	res, err := nm.insertNote.Exec(name, time.Time(datetime).Unix(), text)
	if err != nil {
		return nil, errCreateNote("error with query to db: " + err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, errCreateNote("cannot get id of new note: " + err.Error())
	}

	return &models.Note{
		ID:       id,
		Datetime: &datetime,
		Name:     &name,
		Text:     text,
	}, nil
}

func (nm *NoteMapper) DeleteNote(id int64) error {
	result, err := nm.deleteNote.Exec(id)
	if err != nil {
		return fmt.Errorf("cannot delete note: %s", err)
	}

	n, _ := result.RowsAffected()
	if n == 0 {
		return mappers.ErrFindNoteById
	}

	return nil
}

type errUpdateNote string

func (e errUpdateNote) Error() string {
	return "cannot update note: " + string(e)
}

func (nm *NoteMapper) UpdateNote(id int64, datetime strfmt.DateTime, name, text string) error {
	result, err := nm.updateNoteById.Exec(name, time.Time(datetime), text, id)
	if err != nil {
		return errUpdateNote(err.Error())
	}

	n, _ := result.RowsAffected()
	if n == 0 {
		return mappers.ErrFindNoteById
	}

	return nil
}

func (nm *NoteMapper) GetNoteById(id int64) (*models.Note, error) {
	var (
		note = &models.Note{}

		name          string
		unixtimestamp int64
	)

	err := nm.selectNoteById.QueryRow(id).Scan(&note.ID, &unixtimestamp, &name, &note.Text)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, mappers.ErrFindNoteById
		}
		return nil, fmt.Errorf("cannot get note by id: %s", err)
	}

	note.Name = &name
	var datetime = strfmt.DateTime(time.Unix(unixtimestamp, 0))
	note.Datetime = &datetime

	return note, nil
}

type errFindNotes string

func (e errFindNotes) Error() string {
	return "cannot find notes: "
}

func (nm *NoteMapper) GetNotesByTimeRange(from strfmt.Date, to strfmt.Date) ([]*models.Note, error) {
	if time.Time(from).Unix() >= time.Time(to).Unix() {
		return nil, errFindNotes("invalid time range")
	}

	rows, err := nm.selectNotesByTimeRange.Query(time.Time(from).Unix(), time.Time(to).Unix())
	if err != nil {
		return nil, errFindNotes("cannot get notes by time range: " + err.Error())
	}
	defer rows.Close()

	var notes = make([]*models.Note, 0)

	for rows.Next() {
		var (
			note          = &models.Note{}
			unixtimestamp int64
			name          string
		)
		if err := rows.Scan(&note.ID, &unixtimestamp, &name, &note.Text); err != nil {
			return nil, errFindNotes("cannot get notes by time range: " + err.Error())
		}
		note.Name = &name
		var datetime = strfmt.DateTime(time.Unix(unixtimestamp, 0))
		note.Datetime = &datetime
		notes = append(notes, note)
	}
	if err := rows.Err(); err != nil {
		return nil, errFindNotes("cannot get notes by time range: " + err.Error())
	}

	return notes, nil
}

func (nm *NoteMapper) GetNotesByTimeRangeGrep(from strfmt.Date, to strfmt.Date, name string) ([]*models.Note, error) {
	if len(name) == 0 {
		return nm.GetNotesByTimeRange(from, to)
	}

	if time.Time(from).Unix() >= time.Time(to).Unix() {
		return nil, errFindNotes("invalid time range")
	}

	rows, err := nm.selectNotesByTimeRangeGrep.Query(time.Time(from).Unix(), time.Time(to).Unix(), name)
	if err != nil {
		return nil, errFindNotes("cannot get notes by time range: " + err.Error())
	}
	defer rows.Close()

	var notes = make([]*models.Note, 0)

	for rows.Next() {
		var (
			note          = &models.Note{}
			unixtimestamp int64
			name          string
		)
		if err := rows.Scan(&note.ID, &unixtimestamp, &name, &note.Text); err != nil {
			return nil, errFindNotes("cannot get notes by time range: " + err.Error())
		}
		note.Name = &name
		var datetime = strfmt.DateTime(time.Unix(unixtimestamp, 0))
		note.Datetime = &datetime
		notes = append(notes, note)
	}
	if err := rows.Err(); err != nil {
		return nil, errFindNotes("cannot get notes by time range: " + err.Error())
	}

	return notes, nil
}
