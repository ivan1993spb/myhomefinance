package main

import (
	"database/sql"
	"flag"
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
