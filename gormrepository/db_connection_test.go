package gormrepository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func openDB() (*gorm.DB, error) {
	return gorm.Open("sqlite3", "/tmp/test.db")
}
