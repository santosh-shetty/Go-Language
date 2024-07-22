package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func DBConnect() {
	database, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Erorr During MySQL Connection", err)
	}
	// defer db.Close()
	db = database
}

func GetDB() *gorm.DB {
	return db
}
