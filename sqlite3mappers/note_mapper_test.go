package sqlite3mappers

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ivan1993spb/myhomefinance/mappers"
)

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

	note1, err := noteMapper.CreateNote(time.Unix(1, 0), "test name 1", "test text")
	require.Nil(t, err)
	require.Equal(t, int64(1), note1.Id, "first note cotains invalid id")

	note2, err := noteMapper.CreateNote(time.Unix(2, 0), "test name 2", "test text")
	require.Nil(t, err)
	require.Equal(t, int64(2), note2.Id, "second note cotains invalid id")

	// Try to create note with empty name
	_, err = noteMapper.CreateNote(time.Now(), "", "test text")
	require.NotNil(t, err)

	// Try to get note by id that cannot be found
	_, err = noteMapper.GetNoteById(note2.Id + 100)
	require.Equal(t, mappers.ErrFindNoteById, err)

	// Try to get note by correct id
	myNote, err := noteMapper.GetNoteById(note2.Id)
	require.Nil(t, err)
	// Check if received correct note
	require.Equal(t, note2.Name, myNote.Name)

	// Try to get note list by invalid time range
	_, err = noteMapper.GetNotesByTimeRange(time.Unix(3, 0), time.Unix(3, 0))
	require.NotNil(t, err)

	// Try to get note list by invalid time range
	_, err = noteMapper.GetNotesByTimeRange(time.Unix(3, 0), time.Unix(1, 0))
	require.NotNil(t, err)

	// Try to get note list by correct time range
	notes, err := noteMapper.GetNotesByTimeRange(time.Unix(1, 0), time.Unix(3, 0))
	require.Nil(t, err)
	require.Equal(t, 2, len(notes))

	// Try to delete note by invalid id
	err = noteMapper.DeleteNote(note2.Id + 100)
	require.NotNil(t, err)

	// Try to delete note by valid id
	err = noteMapper.DeleteNote(note1.Id)
	require.Nil(t, err)

	// Try to delete note by valid id
	err = noteMapper.DeleteNote(note2.Id)
	require.Nil(t, err)

	notes, err = noteMapper.GetNotesByTimeRange(time.Unix(1, 0), time.Unix(3, 0))
	require.Nil(t, err)
	require.Equal(t, 0, len(notes))

	notes, err = noteMapper.GetNotesByTimeRangeMatch(time.Unix(1, 0), time.Unix(3, 0), "test")
	require.Nil(t, err)
	require.Equal(t, 0, len(notes))

	// Create notes to test note selectors
	_, err = noteMapper.CreateNote(time.Unix(3, 0), "test name 1", "test text 1")
	require.Nil(t, err)
	_, err = noteMapper.CreateNote(time.Unix(4, 0), "test name 2", "test text 2")
	require.Nil(t, err)
	_, err = noteMapper.CreateNote(time.Unix(5, 0), "test name 3", "test text 3")
	require.Nil(t, err)
	_, err = noteMapper.CreateNote(time.Unix(6, 0), "test name 4", "test text 4")
	require.Nil(t, err)

	// Select notes by first time range
	notes, err = noteMapper.GetNotesByTimeRange(time.Unix(1, 0), time.Unix(3, 0))
	require.Nil(t, err)
	require.Equal(t, 1, len(notes))

	// Test grep selector for case sensitivity
	notes, err = noteMapper.GetNotesByTimeRangeMatch(time.Unix(1, 0), time.Unix(3, 0), "tEst")
	require.Nil(t, err)
	require.Equal(t, 1, len(notes))

	notes, err = noteMapper.GetNotesByTimeRangeMatch(time.Unix(1, 0), time.Unix(7, 0), "not")
	require.Nil(t, err)
	require.Equal(t, 0, len(notes))

	notes, err = noteMapper.GetNotesByTimeRangeMatch(time.Unix(1, 0), time.Unix(7, 0), "te me")
	require.Nil(t, err)
	require.Equal(t, 4, len(notes))
}
