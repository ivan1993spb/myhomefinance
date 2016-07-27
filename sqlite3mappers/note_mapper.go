package sqlite3mappers

import (
	"database/sql"
	"time"

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
	_SQL_SELECT_NOTES_BY_TIME_RANGE_MATCH = "SELECT `id`, `unixtimestamp`, `name`, `text` FROM `notes` " +
		"WHERE `unixtimestamp` BETWEEN ? AND ? AND match(`name`, ?)"
)

type NoteMapper struct {
	db *sql.DB

	// Prepared statements

	insertNote *sql.Stmt
	deleteNote *sql.Stmt

	selectNoteById *sql.Stmt
	updateNoteById *sql.Stmt

	selectNotesByTimeRange      *sql.Stmt
	selectNotesByTimeRangeMatch *sql.Stmt
}

func NewNoteMapper(db *sql.DB) (*NoteMapper, error) {
	err := db.Ping()
	if err != nil {
		return nil, mappers.ErrCreateNoteMapper(err.Error())
	}

	noteMapper := &NoteMapper{db: db}

	noteMapper.insertNote, err = db.Prepare(_SQL_INSERT_NOTE)
	if err != nil {
		return nil, mappers.ErrCreateNoteMapper(err.Error())
	}

	noteMapper.deleteNote, err = db.Prepare(_SQL_DELETE_NOTE_BY_ID)
	if err != nil {
		return nil, mappers.ErrCreateNoteMapper(err.Error())
	}

	noteMapper.selectNoteById, err = db.Prepare(_SQL_SELECT_NOTE_BY_ID)
	if err != nil {
		return nil, mappers.ErrCreateNoteMapper(err.Error())
	}

	noteMapper.updateNoteById, err = db.Prepare(_SQL_UPDATE_NOTE_BY_ID)
	if err != nil {
		return nil, mappers.ErrCreateNoteMapper(err.Error())
	}

	noteMapper.selectNotesByTimeRange, err = db.Prepare(_SQL_SELECT_NOTES_BY_TIME_RANGE)
	if err != nil {
		return nil, mappers.ErrCreateNoteMapper(err.Error())
	}

	noteMapper.selectNotesByTimeRangeMatch, err = db.Prepare(_SQL_SELECT_NOTES_BY_TIME_RANGE_MATCH)
	if err != nil {
		return nil, mappers.ErrCreateNoteMapper(err.Error())
	}

	return noteMapper, nil
}

func (nm *NoteMapper) CreateNote(time time.Time, name string, text string) (*models.Note, error) {
	if len(name) == 0 {
		return nil, mappers.ErrCreateNoteEmptyName
	}

	res, err := nm.insertNote.Exec(name, time.Unix(), text)
	if err != nil {
		return nil, mappers.ErrCreateNote("error with query to db: " + err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, mappers.ErrCreateNote("cannot get id of new note: " + err.Error())
	}

	return &models.Note{
		Id:   id,
		Time: time,
		Name: name,
		Text: text,
	}, nil
}

func (nm *NoteMapper) DeleteNote(id int64) error {
	result, err := nm.deleteNote.Exec(id)
	if err != nil {
		return mappers.ErrDeleteNote(err.Error())
	}

	n, _ := result.RowsAffected()
	if n == 0 {
		return mappers.ErrDeleteNoteNotFound
	}

	return nil
}

func (nm *NoteMapper) UpdateNote(id int64, time time.Time, name, text string) error {
	if len(name) == 0 {
		return mappers.ErrUpdateNoteEmptyName
	}

	result, err := nm.updateNoteById.Exec(name, time.Unix(), text, id)
	if err != nil {
		return mappers.ErrUpdateNote(err.Error())
	}

	n, _ := result.RowsAffected()
	if n == 0 {
		return mappers.ErrUpdateNoteNotFound
	}

	return nil
}

func (nm *NoteMapper) GetNoteById(id int64) (*models.Note, error) {
	var (
		note          = &models.Note{}
		unixtimestamp int64
	)

	err := nm.selectNoteById.QueryRow(id).Scan(&note.Id, &unixtimestamp, &note.Name, &note.Text)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, mappers.ErrGetNoteByIdNotFound
		}
		return nil, mappers.ErrGetNoteById(err.Error())
	}

	note.Time = time.Unix(unixtimestamp, 0)

	return note, nil
}

func (nm *NoteMapper) GetNotesByTimeRange(from time.Time, to time.Time) ([]*models.Note, error) {
	if from.Unix() >= to.Unix() {
		return nil, mappers.ErrGetNotesByTimeRangeInvalidTimeRange
	}

	rows, err := nm.selectNotesByTimeRange.Query(from.Unix(), to.Unix())
	if err != nil {
		return nil, mappers.ErrGetNotesByTimeRange(err.Error())
	}
	defer rows.Close()

	var notes = make([]*models.Note, 0)

	for rows.Next() {
		var (
			note          = &models.Note{}
			unixtimestamp int64
		)
		if err := rows.Scan(&note.Id, &unixtimestamp, &note.Name, &note.Text); err != nil {
			return nil, mappers.ErrGetNotesByTimeRange(err.Error())
		}
		note.Time = time.Unix(unixtimestamp, 0)
		notes = append(notes, note)
	}
	if err := rows.Err(); err != nil {
		return nil, mappers.ErrGetNotesByTimeRange(err.Error())
	}

	return notes, nil
}

func (nm *NoteMapper) GetNotesByTimeRangeMatch(from time.Time, to time.Time, name string) ([]*models.Note, error) {
	if len(name) == 0 {
		return nil, mappers.ErrGetNotesByTimeRangeMatchEmptyName
	}

	if from.Unix() >= to.Unix() {
		return nil, mappers.ErrGetNotesByTimeRangeMatchInvalidTimeRange
	}

	rows, err := nm.selectNotesByTimeRangeMatch.Query(from.Unix(), to.Unix(), name)
	if err != nil {
		return nil, mappers.ErrGetNotesByTimeRangeMatch(err.Error())
	}
	defer rows.Close()

	var notes = make([]*models.Note, 0)

	for rows.Next() {
		var (
			note          = &models.Note{}
			unixtimestamp int64
		)
		if err := rows.Scan(&note.Id, &unixtimestamp, &note.Name, &note.Text); err != nil {
			return nil, mappers.ErrGetNotesByTimeRangeMatch(err.Error())
		}
		note.Time = time.Unix(unixtimestamp, 0)
		notes = append(notes, note)
	}
	if err := rows.Err(); err != nil {
		return nil, mappers.ErrGetNotesByTimeRangeMatch(err.Error())
	}

	return notes, nil
}
