package models

import (
	"errors"
	"fmt"
)

// Magazine struct
type Magazine struct {
	Title       string
	ISBN        string
	Authors     []string
	PublishedAt string
}

// NewMagazine constructor for Magazine types
func NewMagazine(title, isbn, publishedAt string, authors []string) Magazine {
	return Magazine{
		Title:       title,
		ISBN:        isbn,
		Authors:     authors,
		PublishedAt: publishedAt,
	}
}

// PrintProduct method prints data for a magazine in the console
func (m Magazine) PrintProduct() error {
	if m.ISBN != "" {
		fmt.Printf("Magazine Found!\n")
		fmt.Printf("ISBN: %v\n", m.ISBN)
		fmt.Printf("Title: %v\n", m.Title)
		fmt.Print("Authors: ")
		for _, auth := range m.Authors {
			fmt.Print(auth)
		}
		fmt.Println()
		fmt.Printf("Published At: %v\n", m.PublishedAt)
		fmt.Println("======")
	} else {
		return errors.New("product not found")
	}
	return nil
}
