package main

import (
	"database/sql"
	"flag"
	"io/ioutil"
	"os"

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

type Document interface {
}

type Inflow struct {
}

type Outflow struct {
}
