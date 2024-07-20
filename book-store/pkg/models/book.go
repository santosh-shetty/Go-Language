package models

import (
	"github.com/jinzhu/gorm"
	"github.com/santosh/book-store/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// ========== Function Which called by Controller
// GetAllBook retrieves all books from the database.
func GetAllBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// GetBookById retrieves a book by its ID.
func GetBookById(id int64) Book {
	var book Book
	db.First(&book, "id=?", id)
	return book
}

// CreateBook saves a new book to the database.
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// DeleteBook deletes a book by its ID.
func DeleteBook(id int64) error {
	var book Book
	if err := db.Where("id=?", id).Delete(&book).Error; err != nil {
		return err
	}
	return nil
}

// UpdateBook updates an existing book in the database.
func (b *Book) UpdateBook() *Book {
	db.Save(&b)
	return b
}
