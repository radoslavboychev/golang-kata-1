package main

import (
	"log"
	"os"

	"github.com/echocat/golang-kata-1/v1/librarian"
	"github.com/echocat/golang-kata-1/v1/pkg/reader"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	err := godotenv.Load("../.././config/config.env")
	if err != nil {
		log.Println(err)
		return
	}

	// ENV
	magazinesFile := os.Getenv("MAGAZINES_FILE")
	booksFile := os.Getenv("BOOKS_FILE")
	authorEmail := os.Getenv("AUTHOR_EMAIL")
	isbn := os.Getenv("ISBN")
	findTitle := os.Getenv("FIND_TITLE")

	// Load books from file
	books, err := reader.LoadBooks(magazinesFile)
	if err != nil {
		log.Println(err)
		return
	}

	// Load magazines from file
	magazines, err := reader.LoadMagazines(booksFile)
	if err != nil {
		log.Println(err)
		return
	}

	// creates a new products manager (librarian)
	manager := librarian.NewLibrarian(books, magazines)

	// find by ISBN
	p, err := manager.FindByISBN(isbn)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	p.PrintProduct()

	// search for a product by title
	_, err = manager.FindByTitle(findTitle)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// search product by author
	_, err = manager.FindBookByAuthor(authorEmail)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// finding magazine by author
	_, err = manager.FindMagazineByAuthor(authorEmail)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

}
