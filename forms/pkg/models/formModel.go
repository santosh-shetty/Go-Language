package models

import (
	"github.com/jinzhu/gorm"
	"github.com/santosh/forms/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB

// Connection to the Database
func init() {
	config.DBConnect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
}

// ====== Function For Perform action in DB

func (user *User) RegisterUser() *User {
	db.NewRecord(user)
	db.Create(&user)
	return user
}

func IsAuthenticated(email, password string) bool {
	var user User
	db.Where("email=?", email).First(&user)
	if user.ID == 0 {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return false
	}
	return true
}
