package models

import (
	"fmt"
	"math/rand"

	liberror "git.vegaitsourcing.rs/radoslav.boychev/librarian-project/errors"
)

// Book
type Book struct {
	Title       string
	ISBN        string
	Authors     []string
	Description string
}

// NewBook constructor for books
func NewBook(title, isbn, description string, authors []string) Book {
	if len(isbn) < 14 {
		isbn = generateISBN()
	}

	if title == "" {
		title = "Unnamed Book"
	}

	if len(authors) == 0 {
		authors = append(authors, "Unknown Author")
	}

	return Book{
		Title:       title,
		ISBN:        isbn,
		Authors:     authors,
		Description: description,
	}
}

// PrintProduct outputs a string of product information in the console
func (b Book) PrintProduct() error {
	if b.ISBN != "" {
		fmt.Printf("Book: \n")
		fmt.Printf("ISBN: %v\n", b.ISBN)
		fmt.Printf("Title: %v\n", b.Title)
		fmt.Print("Authors: ")
		for _, auth := range b.Authors {
			fmt.Print(auth)
		}
		fmt.Println()
		fmt.Printf("Description: %v\n", b.Description)
		fmt.Println("======")
	} else {
		return liberror.ErrorISBNIsNull
	}
	return nil
}

// generateISBN returns a new ISBN-format string of length 14 divided by the dash '-' symbol
func generateISBN() string {

	min := 1000
	max := 9999

	first := rand.Intn((max - min + 1) + min)
	second := rand.Intn((max - min + 1) + min)
	third := rand.Intn((max - min + 1) + min)

	return fmt.Sprintf("%v-%v-%v", first, second, third)
}
