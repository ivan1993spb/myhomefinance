package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
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

	historyMapper := &HistoryMapper{db}
	historyRecords, err := historyMapper.GetHistoryFeed(time.Unix(0, 0), time.Unix(0, 100))
	require.Nil(t, err, "cannot get history feed")

	var checkBalance float64 = 0 // Calculate balance for each transaction

	for i, hr := range historyRecords {
		checkBalance += hr.Amount
		require.True(t, floatEquals(checkBalance, hr.Balance),
			fmt.Sprintf("on transaction %d: %f != %f", i+1, checkBalance, hr.Balance))
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

const (
	RAW_INFLOW_COUNT  = 100
	RAW_OUTFLOW_COUNT = 100
)

func TestHistoryMapperGetHistoryFeed(t *testing.T) {
	db, err := InitSQLiteDB(TEST_DB_FILE_NAME)
	require.Nil(t, err, "init db returns error")
	defer func() {
		db.Close()
		os.Remove(TEST_DB_FILE_NAME)
	}()

	rand.Seed(time.Now().Unix())

	var start time.Time

	inflowMapper := &InflowMapper{db}

	outflowMapper := &OutflowMapper{db}

	rawInflows := make([]*rawInflow, RAW_INFLOW_COUNT)

	for i := 0; i < RAW_INFLOW_COUNT; i++ {
		rawInflows[i] = &rawInflow{
			time.Unix(int64(rand.Intn(RAW_INFLOW_COUNT)), 0),
			fmt.Sprintf("inflow %d", i),
			rand.Float64() * 100,
		}
	}

	start = time.Now()
	for n, rawInflow := range rawInflows {
		_, err := inflowMapper.CreateInflow(rawInflow.Time, rawInflow.Name, rawInflow.Amount,
			"any desc", "any src")
		require.Nil(t, err, fmt.Sprintf("cannot create inflow #%d with name %q", n, rawInflow.Name))
	}
	t.Logf("creating %d inflow(s) time: %s\n", len(rawInflows), time.Since(start))

	rawOutflows := make([]*rawOutflow, RAW_OUTFLOW_COUNT)

	for i := 0; i < RAW_OUTFLOW_COUNT; i++ {
		rawOutflows[i] = &rawOutflow{
			time.Unix(int64(rand.Intn(RAW_OUTFLOW_COUNT)), 0),
			fmt.Sprintf("outflow %d", i),
			rand.Float64() * 100,
			rand.Float64() * 100,
			rand.Float32(),
		}
	}

	start = time.Now()
	for n, rawOutflow := range rawOutflows {
		_, err := outflowMapper.CreateOutflow(rawOutflow.Time, rawOutflow.Name, rawOutflow.Amount,
			"any desc", "any dst", "any target", 1.5, "any metric unit", 1.0)
		require.Nil(t, err, fmt.Sprintf("cannot create outflow #%d with name %q", n, rawOutflow.Name))
	}
	t.Logf("creating %d outflow(s) time: %s\n", len(rawOutflows), time.Since(start))

	historyMapper := &HistoryMapper{db}

	_, err = historyMapper.GetHistoryFeed(time.Unix(2, 0), time.Unix(0, 0))
	require.NotNil(t, err)

	start = time.Now()
	hf, err := historyMapper.GetHistoryFeed(time.Unix(50, 0), time.Unix(RAW_OUTFLOW_COUNT, 0))
	t.Logf("getting all history (%d row(s)) time: %s\n", len(hf), time.Since(start))
	require.Nil(t, err, "cannot get history feed")
	require.NotEmpty(t, hf, "history feed is empty")

	start = time.Now()
	hf, err = historyMapper.GetHistoryFeed(time.Unix(RAW_OUTFLOW_COUNT+100, 0), time.Unix(RAW_OUTFLOW_COUNT+200, 0))
	t.Log("getting empty history time:", time.Since(start))
	require.Nil(t, err, "cannot get history feed")
	require.Empty(t, hf, "history feed is empty")
}
