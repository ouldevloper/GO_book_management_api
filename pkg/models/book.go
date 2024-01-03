package models

import (
	"restapi/pkg/config"

	"gorm.io/gorm"
	"honnef.co/go/tools/config"
)

type Book struct {
	gorm.Model
	Name        string `gorm;json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&Book{})
}
