package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnDB() (error, *gorm.DB) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:dev123@tcp(103.28.23.185:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, nil
	}
	return nil, db
}
