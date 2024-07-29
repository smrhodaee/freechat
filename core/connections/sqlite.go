package connections

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

func NewSqlite(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	return db, err
}
