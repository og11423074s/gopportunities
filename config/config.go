package config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {

	var err error

	// Initialize SQLite
	db, err = InitSqlite()
	if err != nil {
		return fmt.Errorf("error initializing DB: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	// Init logger
	logger = NewLogger(prefix)
	return logger
}
