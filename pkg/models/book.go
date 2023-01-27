package models

import (
	"errors"
	"fmt"
)

type Book struct {
	Title       string
	ISBN        string
	Authors     []string
	Description string
}

func NewBook(title, isbn, description string, authors []string) Book {
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
		fmt.Printf("Book Found!\n")
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
		return errors.New("product not found")
	}
	return nil
}
