package sqlite3mappers

import (
	"os"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"
)

const TEST_DB_FILE_NAME = "test.db"

func TestNoteMapper(t *testing.T) {
	os.Remove(TEST_DB_FILE_NAME)
	db, err := InitSQLiteDB(TEST_DB_FILE_NAME)
	require.Nil(t, err, "init db returns error")
	defer func() {
		db.Close()
		os.Remove(TEST_DB_FILE_NAME)
	}()

	noteMapper, err := NewNoteMapper(db)
	require.Nil(t, err)

	note, err := noteMapper.CreateNote(strfmt.DateTime(time.Unix(2, 0)), "test name", "test text")
	require.Nil(t, err)
	require.Equal(t, int64(1), note.ID, "note cotains invalid id")

	_, err = noteMapper.CreateNote(strfmt.DateTime(time.Now()), "", "test text")
	require.NotNil(t, err)

	_, err = noteMapper.GetNoteById(note.ID + 1)
	require.NotNil(t, err)

	_, err = noteMapper.GetNoteById(note.ID)
	require.Nil(t, err)

	_, err = noteMapper.GetNotesByTimeRange(strfmt.Date(time.Unix(3, 0)), strfmt.Date(time.Unix(3, 0)))
	require.NotNil(t, err)

	_, err = noteMapper.GetNotesByTimeRange(strfmt.Date(time.Unix(3, 0)), strfmt.Date(time.Unix(1, 0)))
	require.NotNil(t, err)

	notes, err := noteMapper.GetNotesByTimeRange(strfmt.Date(time.Unix(1, 0)), strfmt.Date(time.Unix(3, 0)))
	require.Nil(t, err)
	require.Equal(t, 1, len(notes))

	err = noteMapper.DeleteNote(note.ID + 1)
	require.NotNil(t, err)

	err = noteMapper.DeleteNote(note.ID)
	require.Nil(t, err)

	notes, err = noteMapper.GetNotesByTimeRange(strfmt.Date(time.Unix(1, 0)), strfmt.Date(time.Unix(3, 0)))
	require.Nil(t, err)
	require.Equal(t, 0, len(notes))

	notes, err = noteMapper.GetNotesByTimeRangeGrep(strfmt.Date(time.Unix(1, 0)), strfmt.Date(time.Unix(3, 0)), "test")
	require.Nil(t, err)
	require.Equal(t, 0, len(notes))

	note, err = noteMapper.CreateNote(strfmt.DateTime(time.Unix(3, 0)), "test name 1", "test text 1")
	require.Nil(t, err)

	note, err = noteMapper.CreateNote(strfmt.DateTime(time.Unix(4, 0)), "test name 2", "test text 2")
	require.Nil(t, err)

	note, err = noteMapper.CreateNote(strfmt.DateTime(time.Unix(5, 0)), "test name 3", "test text 3")
	require.Nil(t, err)

	note, err = noteMapper.CreateNote(strfmt.DateTime(time.Unix(6, 0)), "test name 4", "test text 4")
	require.Nil(t, err)

	notes, err = noteMapper.GetNotesByTimeRangeGrep(strfmt.Date(time.Unix(1, 0)), strfmt.Date(time.Unix(3, 0)), "tEst 1")
	require.Nil(t, err)
	require.Equal(t, 1, len(notes))

	notes, err = noteMapper.GetNotesByTimeRangeGrep(strfmt.Date(time.Unix(1, 0)), strfmt.Date(time.Unix(7, 0)), "te me")
	require.Nil(t, err)
	require.Equal(t, 4, len(notes))
}
