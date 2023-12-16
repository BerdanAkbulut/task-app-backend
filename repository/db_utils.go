package repository

import (
	"github.com/BerdanAkbulut/task-app-backend/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func LoadDBInstance() {
	c, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:password@tcp(127.0.0.1:3306)/task?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize:         256,  // default size for string fields
		DisableDatetimePrecision:  true, // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		panic("Couldn't initialize database! Error: " + err.Error())
	}
	db = c

	db.AutoMigrate(&entity.Task{})
	db.AutoMigrate(&entity.User{})
}

func DB() *gorm.DB {
	return db
}
