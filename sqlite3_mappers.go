package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const DEFAULT_SQLITE_DB_FILE_NAME = "myhomefinance.db"

var dbFileName string = DEFAULT_SQLITE_DB_FILE_NAME

func init() {
	flag.StringVar(&dbFileName, "sqlite-db", DEFAULT_SQLITE_DB_FILE_NAME, "sqlite db file name")
}

type ErrSQLiteDB string

func (e ErrSQLiteDB) Error() string {
	return "sqlite db error: " + string(e)
}

// InitSQLiteDB tries to load sqlite db from file or creates new db file with tables and views
func InitSQLiteDB(dbFileName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		return db, ErrSQLiteDB(err.Error())
	}

	if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
		// Try to open file with queries for tables and views creation and execute that queries
		f, err := os.Open("sqlite3_db.sql")
		if err != nil {
			return db, ErrSQLiteDB("cannot init db: " + err.Error())
		}
		defer f.Close()

		rawQuery, err := ioutil.ReadAll(f)
		if err != nil {
			return db, ErrSQLiteDB("cannot read file with queries for tables and views creation: " +
				err.Error())
		}

		_, err = db.Exec(string(rawQuery))
		if err != nil {
			return db, ErrSQLiteDB("cannot execute queries for tables and views creation: " +
				err.Error())
		}
	}

	return db, err
}

type ErrInflowMapper string

func (e ErrInflowMapper) Error() string {
	return "inflow mapper error: " + string(e)
}

type InflowMapper struct {
	*sql.DB
}

// CreateInflow creates new inflow document into db and returns it
func (im *InflowMapper) CreateInflow(t time.Time, name string, amount float64, description, source string,
) (*Inflow, error) {
	if len(name) == 0 {
		return nil, ErrInflowMapper("name cannot be empty")
	}
	if amount <= 0 {
		return nil, ErrInflowMapper(fmt.Sprintf("invalid amount %d (must be > 0)", amount))
	}

	guid, err := newGUID()
	if err != nil {
		return nil, ErrInflowMapper("cannot generate guid: " + err.Error())
	}

	res, err := im.DB.Exec("INSERT INTO `inflow` (`document_guid`, `unixtimestamp`, `name`, `amount`, "+
		"`description`, `source`) VALUES(?, ?, ?, ?, ?, ?)", guid, t.Unix(), name, amount, description,
		source)
	if err != nil {
		return nil, ErrInflowMapper("cannot insert new inflow into db: " + err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, ErrInflowMapper("cannot get id of new inflow: " + err.Error())
	}

	return &Inflow{uint64(id), guid, t, name, amount, description, source}, nil
}

type ErrOutflowMapper string

func (e ErrOutflowMapper) Error() string {
	return "outflow mapper error: " + string(e)
}

type OutflowMapper struct {
	*sql.DB
}

// CreateOutflow creates new outflow document into db and returns it
func (om *OutflowMapper) CreateOutflow(t time.Time, name string, amount float64, description, destination,
	target string, count float64, metricUnit string, satisfaction float32) (*Outflow, error) {
	if len(name) == 0 {
		return nil, ErrOutflowMapper("name cannot be empty")
	}
	if amount <= 0 {
		return nil, ErrOutflowMapper(fmt.Sprintf("invalid amount %d (must be > 0)", amount))
	}

	guid, err := newGUID()
	if err != nil {
		return nil, ErrOutflowMapper("cannot generate guid: " + err.Error())
	}

	res, err := om.DB.Exec("INSERT INTO `outflow` (`document_guid`, `unixtimestamp`, `name`, `amount`, "+
		"`description`, `destination`, `target`, `count`, `metric_unit`, `satisfaction`) VALUES "+
		"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		guid, t.Unix(), name, amount, description, description, target, count, metricUnit, satisfaction)
	if err != nil {
		return nil, ErrOutflowMapper("cannot insert new outflow into db: " + err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, ErrOutflowMapper("cannot get id of new outflow: " + err.Error())
	}

	return &Outflow{uint64(id), guid, t, name, amount, description, destination, target, count, metricUnit,
		satisfaction}, nil
}

type ErrNoteMapper string

func (e ErrNoteMapper) Error() string {
	return "note mapper error: " + string(e)
}

type NoteMapper struct {
	*sql.DB
}

func (nm *NoteMapper) CreateNote(t time.Time, name, text string) (*Note, error) {
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

	return &Note{uint64(id), t, name, text}, nil
}

type ErrHistoryMapper string

func (e ErrHistoryMapper) Error() string {
	return "history mapper error:" + string(e)
}

type HistoryMapper struct {
	*sql.DB
}

func (hm *HistoryMapper) GetHistoryFeed(from, to time.Time) ([]*HistoryRecord, error) {
	if from.Unix() > to.Unix() {
		return nil, ErrHistoryMapper("invalid time range")
	}

	rows, err := hm.DB.Query("SELECT `document_guid`, `unixtimestamp`, `name`, `amount`, `balance` FROM `history` "+
		"WHERE `unixtimestamp` BETWEEN ? AND ?", from.Unix(), to.Unix())
	if err != nil {
		return nil, ErrHistoryMapper("cannot select history records: " + err.Error())
	}
	defer rows.Close()

	historyRecords := make([]*HistoryRecord, 0)

	for rows.Next() {
		var unixtimestamp int64
		hr := &HistoryRecord{}
		if err := rows.Scan(&hr.DocumentGUID, &unixtimestamp, &hr.Name, &hr.Amount, &hr.Balance); err != nil {
			return nil, ErrHistoryMapper("error on scanning query result: " + err.Error())
		}
		hr.Time = time.Unix(unixtimestamp, 0)
		historyRecords = append(historyRecords, hr)
	}

	if err := rows.Err(); err != nil {
		return nil, ErrHistoryMapper("query error: " + err.Error())
	}

	return historyRecords, nil
}