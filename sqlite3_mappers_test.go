package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	TEST_DB_FILE_NAME        = "test.db"
	TEST_DUMP_DATA_FILE_NAME = "testdump.sql"
)

type ErrLoadingSQLQuery string

func (e ErrLoadingSQLQuery) Error() string {
	return "cannot load sql query from file: " + string(e)
}

func loadSQLQuery(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", ErrLoadingSQLQuery("cannot open file: " + err.Error())
	}
	defer f.Close()

	rawQuery, err := ioutil.ReadAll(f)
	if err != nil {
		return "", ErrSQLiteDB("cannot read file: " + err.Error())
	}

	return string(rawQuery), nil
}

var EPSILON float64 = 0.00000001

func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

func TestInitSQLiteDB(t *testing.T) {
	db, err := InitSQLiteDB(TEST_DB_FILE_NAME)
	require.Nil(t, err, "init db return error")
	defer func() {
		db.Close()
		os.Remove(TEST_DB_FILE_NAME)
	}()

	testDumpQuery, err := loadSQLQuery(TEST_DUMP_DATA_FILE_NAME)
	require.Nil(t, err, "load dump data sql query")

	db.Exec(testDumpQuery)

	rows, err := db.Query("SELECT `amount`, `balance` FROM `history` ORDER BY `unixtimestamp` ASC")
	require.Nil(t, err, "error on selecting history query")
	defer rows.Close()

	var (
		transactionNumber int         // Finance transaction number
		checkBalance      float64 = 0 // Calculate balance for each transaction
	)

	for rows.Next() {
		transactionNumber++

		var amount, balance float64
		err := rows.Scan(&amount, &balance)
		require.Nil(t, err, "error on scanning query result")

		checkBalance += amount

		require.True(t, floatEquals(checkBalance, balance),
			fmt.Sprintf("on transaction %d: %f != %f", transactionNumber, checkBalance, balance))
	}

	if err := rows.Err(); err != nil {
		require.Nil(t, err, "error occurred on selecting or scanning rows")
	}
}
