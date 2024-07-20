package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	database, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_todo")
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	db = database
}

func GetDB() *gorm.DB {
	return db
}
