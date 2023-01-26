package models

import (
	"errors"
	"fmt"
)

type Book struct {
	Title       string
	ISBN        string
	Authors     []Author
	Description string
}

func NewBook(title, isbn, description string, authors []Author) Book {
	return Book{
		Title:       title,
		ISBN:        isbn,
		Authors:     authors,
		Description: description,
	}
}

//
func (b Book) PrintProduct() error {
	if b.ISBN != "" {
		fmt.Printf("Book Found!\n")
		fmt.Printf("ISBN: %v\n", b.ISBN)
		fmt.Printf("Title: %v\n", b.Title)
		for _, v := range b.Authors {
			fmt.Println("Authors: ", v.Email, " ", v.Firstname, " ", v.Lastname)
		}
		fmt.Printf("Description: %v\n", b.Description)
		fmt.Println("======")
	} else {
		fmt.Println("Product not found")
		return errors.New("Product not found")
	}
	return nil
}
