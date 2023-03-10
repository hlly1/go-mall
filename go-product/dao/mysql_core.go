package dao

import (
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var DB *gorm.DB

var err error

func init() {
	db_connect := "root:123456@tcp(127.0.0.1:3306)/go-product?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       db_connect, // data source name
		DefaultStringSize:         256,        // default size for string fields
		DisableDatetimePrecision:  true,       // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,       // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,       // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,      // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}
