package main

import (
	"log"

	"github.com/echocat/golang-kata-1/v1/librarian"
	"github.com/echocat/golang-kata-1/v1/reader"
)

func main() {

	books, err := reader.LoadBooks("./resources/books.csv")
	if err != nil {
		log.Println("Error")
	}

	magazines, err := reader.LoadMagazines("./resources/magazines.csv")
	if err != nil {
		log.Println("Error")
	}

	// creates a new products manager (librarian)
	manager := librarian.NewLibrarian(books, magazines)

	// find by ISBN
	p, err := manager.FindByISBN("4545-8558-3232")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	p.PrintProduct()

	// search for a product by title
	_, err = manager.FindByTitle("Das")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

}
