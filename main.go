package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Book model
type Book struct {
	ID     uint `gorm:"primaryKey"`
	Title  string
	Author string
}

func main() {
	// Connect to the database
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto-migrate the schema
	db.AutoMigrate(&Book{})

	// Create a new book
	db.Create(&Book{Title: "The Go Programming Language", Author: "Alan A. A. Donovan"})

	// Read all books
	var books []Book
	db.Find(&books)
	fmt.Println("All Books:")
	for _, book := range books {
		fmt.Printf("- %s by %s\n", book.Title, book.Author)
	}

	// Update a book
	var book Book
	db.First(&book, 1) // find book with ID 1
	db.Model(&book).Update("Author", "Alan Donovan & Brian Kernighan")

	// Delete a book
	db.Delete(&Book{}, 1)
}
