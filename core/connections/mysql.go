package connections

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMySql(dsn string) (*gorm.DB, error) {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Silent),
        DisableForeignKeyConstraintWhenMigrating: true,
    })
    if err != nil {
        return nil, err
    }
    return db, err
} 
