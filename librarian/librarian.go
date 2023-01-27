package librarian

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/echocat/golang-kata-1/v1/models"
	"github.com/echocat/golang-kata-1/v1/reader"
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

// FindBookByAuthor looks up a book based on its author
func (m *Librarian) FindBookByAuthor(email string) ([]models.Book, error) {

	authors, _ := reader.LoadAuthors("./resources/authors.csv")
	records, err := loadFile("./resources/books.csv")
	if err != nil {
		return nil, err
	}

	res := []models.Book{}
	for _, rec := range records {
		b := models.Book{
			Title:       rec[0],
			ISBN:        rec[1],
			Authors:     []string{rec[2]},
			Description: rec[3],
		}

		res = append(res, b)
	}

	bookSlice := []models.Book{}

	for _, book := range res {
		for _, a := range book.Authors {
			if strings.Contains(a, email) {
				bookSlice = append(bookSlice, book)
			}
		}
	}

	out, err := reader.ResolveBookAuthors(authors, bookSlice)
	if err != nil {
		return nil, err
	}

	for _, v := range out {
		v.PrintProduct()
	}
	return res, nil
}

func (m *Librarian) FindMagazineByAuthor(email string) ([]models.Magazine, error) {
	// READ AUTHORS FILE
	authors, _ := reader.LoadAuthors("./resources/authors.csv")
	// LOAD RAW FILES FOR BOOKS
	records, err := loadFile("./resources/magazines.csv")
	if err != nil {
		return nil, err
	}

	res := []models.Magazine{}
	for _, rec := range records {
		b := models.Magazine{
			Title:       rec[0],
			ISBN:        rec[1],
			Authors:     []string{rec[2]},
			PublishedAt: rec[3],
		}

		res = append(res, b)
	}

	for _, mag := range res {
		for _, a := range mag.Authors {
			if a == email {
				res = append(res, mag)
			}
		}
	}

	// ITERATE THROUGH BOOKS
	// read magazine from file
	records, err = loadFile("./resources/magazines.csv")
	if err != nil {
		return nil, err
	}

	mags := []models.Magazine{}
	for _, rec := range records {
		m := models.Magazine{
			Title:       rec[0],
			ISBN:        rec[1],
			Authors:     []string{rec[2]},
			PublishedAt: rec[3],
		}

		mags = append(mags, m)
	}

	magSlice := []models.Magazine{}

	for _, magazine := range mags {
		for _, a := range magazine.Authors {
			if strings.Contains(a, email) {
				magSlice = append(magSlice, magazine)
			}
		}
	}

	out, err := reader.ResolveMagAuthors(authors, magSlice)
	if err != nil {
		return nil, err
	}

	for _, v := range out {
		v.PrintProduct()
	}
	return out, nil
}

// loadFile read the contents from a file
func loadFile(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'

	if _, err := r.Read(); err != nil {
		return nil, err
	}

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
