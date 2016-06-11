package sqlite3mappers

import (
	"database/sql"
	"os"
	"regexp"

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
    amount        DOUBLE NOT NULL,
    description   TEXT,
    source        VARCHAR(300) NOT NULL
);

CREATE TABLE IF NOT EXISTS outflow (
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    document_guid VARCHAR(36) NOT NULL,
    unixtimestamp INT8 DEFAULT (strftime('%s', 'now')) NOT NULL,
    name          VARCHAR(300) NOT NULL,
    amount        DOUBLE NOT NULL,
    description   TEXT,
    destination   VARCHAR(300) NOT NULL,
    target        VARCHAR(300),
    count         DOUBLE,
    metric_unit   VARCHAR(100),
    satisfaction  FLOAT
);

CREATE VIEW IF NOT EXISTS transactions AS
    SELECT * FROM (
        SELECT document_guid, unixtimestamp, name, amount, description FROM inflow
        UNION
        SELECT document_guid, unixtimestamp, name, amount AS amount, description FROM outflow
    ) result_union
    ORDER BY (result_union.unixtimestamp) DESC;
`

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
		_, err = db.Exec(_TABLES_QUERY)
		if err != nil {
			return db, ErrSQLiteDB("cannot execute queries for tables and views creation: " + err.Error())
		}
	}

	return db, err
}
