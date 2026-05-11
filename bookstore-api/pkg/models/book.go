package models

import(
	"github.com/Parship12/bookstore-api/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect() // Calls the config function to establish the database connection.
	db = config.GetDB() // Gets the database connection from config and stores it in the local db variable.
	db.AutoMigrate(&Book{}) // Auto-creates/updates the books table to match the Book struct.
}

// The (b *Book) part makes this a method on the Book type, not a standalone function.
// func - Function keyword
// (b *Book) - Method receiver (this function belongs to Book)
// CreateBook() - Method name
// *Book - Return type (pointer to Book)
func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // NewRecord(b) method inside gorm - Returns true if the record doesn't exist in the database yet
	db.Create(&b) // db.Create(&b) - GORM method that inserts the record. &b - Passes a pointer to the book. GORM converts this to: INSERT INTO books (name, author, publication) VALUES (...)
	return b
}

// Retrieves all books from the database.
func GetAllBooks() []Book {
	var Books []Book // Creates an empty slice to hold the results
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
