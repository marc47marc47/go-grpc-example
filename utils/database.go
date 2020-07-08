package utils

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	. "github.com/markbest/go-grpc-example/conf"
)

var db *gorm.DB

func DB() *gorm.DB {
	if db == nil {
		newDb, err := newDB()
		if err != nil {
			panic(err)
		}
		newDb.DB().SetMaxIdleConns(10)
		newDb.DB().SetMaxOpenConns(100)
		newDb.LogMode(false)
		db = newDb
	}
	return db
}

func newDB() (*gorm.DB, error) {
	if err := InitConfig(); err != nil {
		log.Fatal(err)
	}

	sqlConnection := Conf.DB.User + ":" + Conf.DB.Password + "@tcp(" + Conf.DB.Host + ":" + Conf.DB.Port + ")/" + Conf.DB.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", sqlConnection)
	if err != nil {
		return nil, err
	}
	return db, nil
}
