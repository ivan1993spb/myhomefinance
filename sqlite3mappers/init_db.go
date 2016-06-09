package sqlite3mappers

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

//go:generate go-bindata --pkg=sqlite3mappers db.sql

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
		_, err = db.Exec(string(MustAsset("db.sql")))
		if err != nil {
			return db, ErrSQLiteDB("cannot execute queries for tables and views creation: " + err.Error())
		}
	}

	return db, err
}
