package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

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
	require.Nil(t, err, "init db returns error")
	defer func() {
		db.Close()
		os.Remove(TEST_DB_FILE_NAME)
	}()

	testDumpQuery, err := loadSQLQuery(TEST_DUMP_DATA_FILE_NAME)
	require.Nil(t, err, "load dump data sql query")

	_, err = db.Exec(testDumpQuery)
	require.Nil(t, err)

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

func TestInflowMapperCreateInflow(t *testing.T) {
	db, err := InitSQLiteDB(TEST_DB_FILE_NAME)
	require.Nil(t, err, "init db returns error")
	defer func() {
		db.Close()
		os.Remove(TEST_DB_FILE_NAME)
	}()

	inflowMapper := &InflowMapper{db}

	var amount float64 = 6.25
	inflow, err := inflowMapper.CreateInflow(time.Now(), "test inflow", amount, "any desc", "any src")
	require.Nil(t, err)
	require.Equal(t, amount, inflow.Amount, "inflow contains invalid amount")
	require.Equal(t, uint64(1), inflow.Id, "inflow contains invalid id")

	_, err = inflowMapper.CreateInflow(time.Now(), "test inflow", 0, "any desc", "any src")
	require.NotNil(t, err)

	_, err = inflowMapper.CreateInflow(time.Now(), "", 1.00, "any desc", "any src")
	require.NotNil(t, err)
}

func TestInflowMapperCreateOutflow(t *testing.T) {
	db, err := InitSQLiteDB(TEST_DB_FILE_NAME)
	require.Nil(t, err, "init db returns error")
	defer func() {
		db.Close()
		os.Remove(TEST_DB_FILE_NAME)
	}()

	outflowMapper := &OutflowMapper{db}

	var amount float64 = 1.55
	outflow, err := outflowMapper.CreateOutflow(time.Now(), "test outflow", amount, "any desc", "any dst", "any target",
		1.5, "kg", 1.0)
	require.Nil(t, err)
	require.Equal(t, amount, outflow.Amount, "outflow contains invalid amount")
	require.Equal(t, uint64(1), outflow.Id, "outflow contains invalid id")

	_, err = outflowMapper.CreateOutflow(time.Now(), "test outflow", 0, "any desc", "any dst", "any target", 1.5, "kg",
		1.0)
	require.NotNil(t, err)

	_, err = outflowMapper.CreateOutflow(time.Now(), "", 1.00, "any desc", "any dst", "any target", 1.5, "kg", 1.0)
	require.NotNil(t, err)
}

func TestNoteMapperCreateNote(t *testing.T) {
	db, err := InitSQLiteDB(TEST_DB_FILE_NAME)
	require.Nil(t, err, "init db returns error")
	defer func() {
		db.Close()
		os.Remove(TEST_DB_FILE_NAME)
	}()

	noteMapper := &NoteMapper{db}

	note, err := noteMapper.CreateNote(time.Now(), "test name", "test text")
	require.Nil(t, err)
	require.Equal(t, uint64(1), note.Id, "note cotains invalid id")

	_, err = noteMapper.CreateNote(time.Now(), "", "test text")
	require.NotNil(t, err)
}

type rawInflow struct {
	Time   time.Time
	Name   string
	Amount float64
}

type rawOutflow struct {
	Time         time.Time
	Name         string
	Amount       float64
	Count        float64
	Satisfaction float32
}

func TestHistoryMapperGetHistoryFeed(t *testing.T) {
	db, err := InitSQLiteDB(TEST_DB_FILE_NAME)
	require.Nil(t, err, "init db returns error")
	defer func() {
		db.Close()
		os.Remove(TEST_DB_FILE_NAME)
	}()

	inflowMapper := &InflowMapper{db}

	outflowMapper := &OutflowMapper{db}

	rawInflows := []*rawInflow{
		{time.Unix(1, 0), "inflow 1", 1.55},
		{time.Unix(2, 0), "inflow 2", 7.50},
		{time.Unix(3, 0), "inflow 3", 1.65},
		{time.Unix(4, 0), "inflow 4", 5.25},
		{time.Unix(5, 0), "inflow 5", 1.00},
		{time.Unix(6, 0), "inflow 6", 0.01},
		{time.Unix(7, 0), "inflow 7", 1.55},
		{time.Unix(8, 0), "inflow 8", 7.65},
		{time.Unix(9, 0), "inflow 9", 10.25},
		{time.Unix(15, 0), "inflow 10", 5.25},
		{time.Unix(16, 0), "inflow 11", 1.25},
		{time.Unix(23, 0), "inflow 12", 2.25},
		{time.Unix(25, 0), "inflow 13", 3.25},
		{time.Unix(26, 0), "inflow 14", 4.25},
	}

	for n, rawInflow := range rawInflows {
		_, err := inflowMapper.CreateInflow(rawInflow.Time, rawInflow.Name, rawInflow.Amount,
			"any desc", "any src")
		require.Nil(t, err, fmt.Sprintf("cannot create inflow #%d with name %q", n, rawInflow.Name))
	}

	rawOutflows := []*rawOutflow{
		{time.Unix(5, 0), "outflow 1", 1.25, 2, 0.5},
		{time.Unix(7, 0), "outflow 2", 0.20, 1, 0.5},
		{time.Unix(8, 0), "outflow 3", 3.00, 5, 0.5},
		{time.Unix(9, 0), "outflow 4", 0.10, 3, 0.5},
		{time.Unix(12, 0), "outflow 5", 5.00, 5, 0.5},
		{time.Unix(23, 0), "outflow 6", 2.80, 1, 0.5},
		{time.Unix(25, 0), "outflow 7", 7.35, 5, 0.5},
		{time.Unix(30, 0), "outflow 8", 6.25, 2, 0.5},
	}

	for n, rawOutflow := range rawOutflows {
		_, err := outflowMapper.CreateOutflow(rawOutflow.Time, rawOutflow.Name, rawOutflow.Amount,
			"any desc", "any dst", "any target", 1.5, "any metric unit", 1.0)
		require.Nil(t, err, fmt.Sprintf("cannot create outflow #%d with name %q", n, rawOutflow.Name))
	}

	historyMapper := &HistoryMapper{db}

	_, err = historyMapper.GetHistoryFeed(time.Unix(2, 0), time.Unix(1, 0))
	require.NotNil(t, err)

	hf, err := historyMapper.GetHistoryFeed(time.Unix(0, 0), time.Unix(100, 0))
	require.Nil(t, err, "cannot get history feed")
	require.NotEmpty(t, hf, "history feed is empty")

	hf, err = historyMapper.GetHistoryFeed(time.Unix(31, 0), time.Unix(100, 0))
	require.Nil(t, err, "cannot get history feed")
	require.Empty(t, hf, "history feed is empty")
}
