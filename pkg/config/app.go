package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func connect() {
	d, err := gorm.Open(mysql.Open("books_api:user@tcp(127.0.0.1:3306)/books_api??charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		// log.Fatal("DB Error : ", err.Error())
		panic(err)
	}
	db = d
}

func getDB() *gorm.DB {
	return db
}
