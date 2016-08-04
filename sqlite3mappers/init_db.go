package sqlite3mappers

import (
	"database/sql"
	"fmt"
	"os"

	sqlite "github.com/mattn/go-sqlite3"
)

const _TABLES_QUERY = `
CREATE TABLE IF NOT EXISTS notes (
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    name          VARCHAR(300) NOT NULL,
    unixtimestamp INT8 DEFAULT (strftime('%s', 'now')) NOT NULL,
    text          TEXT
);

CREATE TABLE IF NOT EXISTS inflow (
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    document_guid VARCHAR(36) NOT NULL,
    unixtimestamp INT8 DEFAULT (strftime('%s', 'now')) NOT NULL,
    name          VARCHAR(300) NOT NULL,
    amount        DECIMAL(10,2) NOT NULL,
    description   TEXT,
    source        VARCHAR(300) NOT NULL
);

CREATE TABLE IF NOT EXISTS outflow (
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    document_guid VARCHAR(36) NOT NULL,
    unixtimestamp INT8 DEFAULT (strftime('%s', 'now')) NOT NULL,
    name          VARCHAR(300) NOT NULL,
    amount        DECIMAL(10,2) NOT NULL,
    description   TEXT,
    destination   VARCHAR(300) NOT NULL,
    target        VARCHAR(300),
    count         DOUBLE,
    metric_unit   VARCHAR(100),
    satisfaction  FLOAT
);
`

func init() {
	sql.Register("sqlite3_mhf", &sqlite.SQLiteDriver{
		ConnectHook: func(conn *sqlite.SQLiteConn) error {
			if err := conn.RegisterFunc("match", match, true); err != nil {
				return err
			}

			return nil
		},
	})
}

// InitSQLiteDB tries to load sqlite db from file or creates new db file with all necessary tables
func InitSQLiteDB(dbFileName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3_mhf", dbFileName)
	if err != nil {
		return db, fmt.Errorf("init sqlite db error: %s", err)
	}

	if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
		_, err = db.Exec(_TABLES_QUERY)
		if err != nil {
			panic("cannot execute init db queries for tables and views creation: " + err.Error())
		}
	}

	return db, err
}
