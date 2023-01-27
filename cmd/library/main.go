package main

import (
	"log"

	"github.com/echocat/golang-kata-1/v1/config"
	"github.com/echocat/golang-kata-1/v1/librarian"
	"github.com/echocat/golang-kata-1/v1/pkg/reader"
)

func main() {

	// Load configuration
	conf, err := config.LoadConfig()
	if err != nil {
		log.Println(err)
		return
	}

	// Load books from file
	books, err := reader.LoadBooks(conf.BooksFile)
	if err != nil {
		log.Println(err)
		return
	}

	// Load magazines from file
	magazines, err := reader.LoadMagazines(conf.MagazinesFile)
	if err != nil {
		log.Println(err)
		return
	}

	// creates a new products manager (librarian)
	manager := librarian.NewLibrarian(books, magazines)

	// find by ISBN
	p, err := manager.FindByISBN(conf.FindByISBN)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	p.PrintProduct()

	// search for a product by title
	_, err = manager.FindByTitle(conf.FindByTitle)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// search product by author
	_, err = manager.FindBookByAuthor(conf.AuthorEmail)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// finding magazine by author
	_, err = manager.FindMagazineByAuthor(conf.AuthorEmail)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// TODO improve resolving authors
	// TODO edge cases

}
