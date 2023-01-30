package librarian

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	liberror "github.com/echocat/golang-kata-1/v1/errors"
	"github.com/echocat/golang-kata-1/v1/pkg/models"
	"github.com/echocat/golang-kata-1/v1/pkg/reader"
)

// Librarian
type Librarian struct {
	Books     []models.Book
	Magazines []models.Magazine
}

// NewLibrarian constructor creates a new product Librarian
func NewLibrarian(books []models.Book, magazines []models.Magazine) *Librarian {
	return &Librarian{
		Books:     books,
		Magazines: magazines,
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
			return liberror.ErrFailedToOpenFile
		}
	}
	return nil
}

// FindByISBN looks up a product by ISBN and returns it
func (m *Librarian) FindByISBN(isbn string) (models.Product, error) {
	var p models.Product

	if len(isbn) != 14 {
		return nil, liberror.ErrFailedToFindProduct
	}

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

	if p == nil {
		return nil, liberror.ErrFailedToFindProduct
	}

	return p, nil
}

// FindByTitle looks up a product by title
func (m *Librarian) FindByTitle(title string) ([]models.Product, error) {

	if title == "" {
		return nil, liberror.ErrFailedToFindProduct
	}

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

	if foundItems == nil {
		return nil, liberror.ErrFailedToFindProduct
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

	filepathAuthors := os.Getenv("AUTHORS_FILE")
	filepathBooks := os.Getenv("BOOKS_FILE")

	authors, _ := reader.LoadAuthors(filepathAuthors)

	records, err := loadFile(filepathBooks)
	if err != nil {
		return nil, liberror.ErrFailedToOpenFile
	}

	if email == "" {
		return nil, liberror.ErrInvalidEmail
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
		return nil, liberror.ErrFailedToResolveAuthors
	}

	for _, v := range out {
		v.PrintProduct()
	}
	return res, nil
}

// FindMagazineByAuthor returns all magazines whose author matches the provided author email
func (m *Librarian) FindMagazineByAuthor(email string) ([]models.Magazine, error) {

	filepathAuthors := os.Getenv("AUTHORS_FILE")
	filepathMagazines := os.Getenv("MAGAZINES_FILE")

	authors, _ := reader.LoadAuthors(filepathAuthors)
	records, err := loadFile(filepathMagazines)
	if err != nil {
		return nil, liberror.ErrFailedToOpenFile
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

	records, err = loadFile(filepathMagazines)
	if err != nil {
		return nil, liberror.ErrFailedToOpenFile
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
		return nil, liberror.ErrFailedToResolveAuthors
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
		return nil, liberror.ErrGeneric
	}

	records, err := r.ReadAll()
	if err != nil {
		return nil, liberror.ErrGeneric
	}

	return records, nil
}
