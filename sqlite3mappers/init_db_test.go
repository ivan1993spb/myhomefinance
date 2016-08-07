package sqlite3mappers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const testDBName = "test.db"

func TestInitSQLiteDB(t *testing.T) {
	os.Remove(testDBName)
	db, err := InitSQLiteDB(testDBName)
	require.Nil(t, err, "init db returns error")
	defer func() {
		db.Close()
		os.Remove(testDBName)
	}()

	var grepTest bool

	err = db.QueryRow(`SELECT match("ok","ok")`).Scan(&grepTest)
	require.Nil(t, err, "select match returns error")
	require.True(t, grepTest)

	err = db.QueryRow(`SELECT match("test tset olol 123","tset  123")`).Scan(&grepTest)
	require.Nil(t, err, "select match returns error")
	require.True(t, grepTest)

	err = db.QueryRow(`SELECT match("test","123")`).Scan(&grepTest)
	require.Nil(t, err, "select match returns error")
	require.False(t, grepTest)
}
