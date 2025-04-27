# GORM SQLite Example

This project demonstrates basic CRUD (Create, Read, Update, Delete) operations using GORM with an SQLite database in Go.

## Prerequisites

*   Go installed (version 1.16 or higher)
*   GORM installed (`go get -u gorm.io/gorm`)
*   SQLite driver for GORM installed (`go get -u gorm.io/driver/sqlite`)

## Code Description

The `main.go` file contains the following:

*   **Package declaration:** `package main` declares the package as the entry point of the application.
*   **Import statements:**
    *   `"fmt"`: Imports the `fmt` package for formatted I/O.
    *   `"gorm.io/driver/sqlite"`: Imports the SQLite driver for GORM.
    *   `"gorm.io/gorm"`: Imports the GORM library.
*   **Book model:**
    ```go
    type Book struct {
        ID     uint `gorm:"primaryKey"`
        Title  string
        Author string
    }
    ```
    Defines the `Book` struct, which represents a book in the database. The `gorm:"primaryKey"` tag specifies that the `ID` field is the primary key.
*   **`main` function:**
    *   **Connect to the database:**
        ```go
        db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
        if err != nil {
            panic("failed to connect database")
        }
        ```
        Connects to the SQLite database file named `books.db`. If the connection fails, the program panics.
    *   **Auto-migrate the schema:**
        ```go
        db.AutoMigrate(&Book{})
        ```
        Automatically creates or updates the `books` table in the database based on the `Book` struct.
    *   **Create a new book:**
        ```go
        db.Create(&Book{Title: "The Go Programming Language", Author: "Alan A. A. Donovan"})
        ```
        Creates a new `Book` record in the database with the specified title and author.
    *   **Read all books:**
        ```go
        var books []Book
        db.Find(&books)
        fmt.Println("All Books:")
        for _, book := range books {
            fmt.Printf("- %s by %s\n", book.Title, book.Author)
        }
        ```
        Retrieves all `Book` records from the database and prints them to the console.
    *   **Update a book:**
        ```go
        var book Book
        db.First(&book, 1) // find book with ID 1
        db.Model(&book).Update("Author", "Alan Donovan & Brian Kernighan")
        ```
        Finds the `Book` record with ID 1 and updates the author.
    *   **Delete a book:**
        ```go
        db.Delete(&Book{}, 1)
        ```
        Deletes the `Book` record with ID 1.

## Usage

1.  **Clone the repository:**

    ```bash
    git clone <repository_url>
    cd <repository_directory>
    ```

2.  **Run the program:**

    ```bash
    go run main.go
    ```

    This will:

    *   Create an SQLite database file named `books.db` (if it doesn't exist).
    *   Create the `books` table (if it doesn't exist).
    *   Insert a new book record.
    *   Retrieve and print all book records.
    *   Update the author of the book with ID 1.
    *   Delete the book with ID 1.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.