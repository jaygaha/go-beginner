package sqlite_db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("../gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
