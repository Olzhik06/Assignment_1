package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func NewLibrary() *Library {
	return &Library{Books: make(map[string]Book)}
}

func (l *Library) AddBook(book Book) {
	l.Books[book.ID] = book
	fmt.Println("Book added successfully!")
}

func (l *Library) BorrowBook(id string) {
	if book, exists := l.Books[id]; exists {
		if !book.IsBorrowed {
			book.IsBorrowed = true
			l.Books[id] = book
			fmt.Println("Book borrowed successfully!")
		} else {
			fmt.Println("Book is already borrowed.")
		}
	} else {
		fmt.Println("Book not found.")
	}
}

func (l *Library) ReturnBook(id string) {
	if book, exists := l.Books[id]; exists {
		if book.IsBorrowed {
			book.IsBorrowed = false
			l.Books[id] = book
			fmt.Println("Book returned successfully!")
		} else {
			fmt.Println("Book was not borrowed.")
		}
	} else {
		fmt.Println("Book not found.")
	}
}

func (l *Library) ListBooks() {
	fmt.Println("Available Books:")
	for _, book := range l.Books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

func main() {
	library := NewLibrary()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Book")
		fmt.Println("2. Borrow Book")
		fmt.Println("3. Return Book")
		fmt.Println("4. List Books")
		fmt.Println("5. Exit")

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			fmt.Println("Enter Book ID:")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())

			fmt.Println("Enter Book Title:")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())

			fmt.Println("Enter Book Author:")
			scanner.Scan()
			author := strings.TrimSpace(scanner.Text())

			library.AddBook(Book{ID: id, Title: title, Author: author, IsBorrowed: false})
		case "2":
			fmt.Println("Enter Book ID to Borrow:")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())
			library.BorrowBook(id)
		case "3":
			fmt.Println("Enter Book ID to Return:")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())
			library.ReturnBook(id)
		case "4":
			library.ListBooks()
		case "5":
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
