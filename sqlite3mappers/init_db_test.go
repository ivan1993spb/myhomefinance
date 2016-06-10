package sqlite3mappers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const TEST_DB_FILE_NAME = "test.db"

func TestInitSQLiteDB(t *testing.T) {
	os.Remove(TEST_DB_FILE_NAME)
	db, err := InitSQLiteDB(TEST_DB_FILE_NAME)
	require.Nil(t, err, "init db returns error")
	defer func() {
		db.Close()
		os.Remove(TEST_DB_FILE_NAME)
	}()

	var grepTest bool
	err = db.QueryRow(`SELECT grep("ok","ok")`).Scan(&grepTest)
	require.Nil(t, err, "grep select returns error")
	require.True(t, grepTest)
}
