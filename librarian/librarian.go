package librarian

import (
	"strings"

	"github.com/echocat/golang-kata-1/v1/models"
)

type Librarian struct {
	Books     []models.Book
	Magazines []models.Magazine
	// Authors   []Author
}

// NewLibrarian constructor creates a new product Librarian
func NewLibrarian(books []models.Book, magazines []models.Magazine) *Librarian {
	return &Librarian{
		Books:     books,
		Magazines: magazines,
		// Authors:   authors,
	}
}

// PrintBooks prints the info for all books from the librarian
func (m Librarian) PrintBooks() {
	for _, bookIn := range m.Books {

		bookIn.PrintProduct()
	}
}

// PrintMagazines prints the info for all the magazines from the librarian
func (m Librarian) PrintMagazines() error {
	for _, magIn := range m.Magazines {

		err := magIn.PrintProduct()
		if err != nil {
			return err
		}
	}
	return nil
}

// FindByISBN looks up a product by ISBN and returns it
func (m *Librarian) FindByISBN(isbn string) (models.Product, error) {
	var p models.Product

	for _, mag := range m.Magazines {
		if mag.ISBN == isbn {
			p = mag
		}
	}

	for _, book := range m.Books {
		if book.ISBN == isbn {
			p = book
		}

	}

	return p, nil
}

// FindByTitle looks up a product by title
func (m *Librarian) FindByTitle(title string) ([]models.Product, error) {

	var foundItems []models.Product

	for _, m := range m.Magazines {
		if strings.Contains(m.Title, title) {
			foundItems = append(foundItems, m)
		}
	}

	for _, b := range m.Books {
		if strings.Contains(b.Title, title) {
			foundItems = append(foundItems, b)
		}
	}

	for _, v := range foundItems {
		v.PrintProduct()
	}

	return foundItems, nil
}

// PrintAll prints info for all products
func (m Librarian) PrintAll() {
	m.PrintBooks()
	m.PrintMagazines()
}
