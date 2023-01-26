package main

import (
	"errors"
	"log"

	"github.com/echocat/golang-kata-1/v1/models"
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

	authors, err := reader.LoadAuthors("./resources/authors.csv")
	if err != nil {
		log.Println("Error")
	}
	mags := reader.ResolveAuthors(authors, magazines)

	manager := models.NewManager(books, mags, authors)

	p, err := manager.FindByISBN("2547-8548-2541")
	if err != nil {
		log.Printf("error %v", errors.New("failed to return product"))
	}
	p.PrintProduct()
}
