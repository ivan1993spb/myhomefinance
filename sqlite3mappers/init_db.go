package sqlite3mappers

import (
	"database/sql"
	"os"
	"regexp"

	sqlite "github.com/mattn/go-sqlite3"
)

//go:generate go-bindata --pkg=sqlite3mappers db.sql

func init() {
	sql.Register("sqlite3_mhf", &sqlite.SQLiteDriver{
		ConnectHook: func(conn *sqlite.SQLiteConn) error {
			if err := conn.RegisterFunc("grep", grep, true); err != nil {
				return err
			}

			return nil
		},
	})
}

var spaceExpr = regexp.MustCompile(`\s+`)

func grep(s1, s2 string) bool {
	grepExpr, err := regexp.Compile("(?i)" + spaceExpr.ReplaceAllLiteralString(regexp.QuoteMeta(s2), ".+"))
	if err != nil {
		return false
	}
	return grepExpr.MatchString(s1)
}

type ErrSQLiteDB string

func (e ErrSQLiteDB) Error() string {
	return "sqlite db error: " + string(e)
}

// InitSQLiteDB tries to load sqlite db from file or creates new db file with tables and views
func InitSQLiteDB(dbFileName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3_mhf", dbFileName)
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
