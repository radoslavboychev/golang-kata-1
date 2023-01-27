package librarian

import (
	"testing"

	"github.com/echocat/golang-kata-1/v1/config"
	liberror "github.com/echocat/golang-kata-1/v1/errors"
	"github.com/echocat/golang-kata-1/v1/pkg/models"
	"github.com/echocat/golang-kata-1/v1/pkg/reader"
	"github.com/stretchr/testify/assert"
)

func TestFindByISBN(t *testing.T) {

	conf, err := config.LoadConfig()
	if err != nil {
		return
	}

	books, _ := reader.LoadBooks(conf.BooksFile)
	magazines, _ := reader.LoadMagazines(conf.BooksFile)

	m := NewLibrarian(books, magazines)

	t.Run("CASE_ISBN_INVALID", func(t *testing.T) {

		// Act
		m.FindByISBN("4545-8558-3232-22948")

		// Assert
		assert.EqualError(t, liberror.ErrISBNInvalid, "ISBN is of invalid length")

	})

	t.Run("CASE_ISBN_VALID_NO_RESULTS", func(t *testing.T) {

		// Act
		m.FindByISBN("1024-1024-1024")

		// Assert
		assert.EqualError(t, liberror.ErrFailedToFindProduct, "failed to find product")

	})

	t.Run("CASE_ISBN_VALID_RESULT_FOUND", func(t *testing.T) {

		// Arrange
		expectedP := models.Book{
			Title:       "Schlank im Schlaf",
			ISBN:        "4545-8558-3232",
			Authors:     []string{"null-gustafsson@echocat.org"},
			Description: "Schlank im Schlaf klingt wie ein schöner Traum,aber es ist wirklich möglich. Allerdings nicht nach einer Salamipizza zum Abendbrot. Die Grundlagen dieses neuartigen Konzepts sind eine typgerechte Insulin-Trennkost sowie Essen und Sport im Takt der biologischen Uhr. Wie die Bio-Uhr tickt und was auf dem Speiseplan stehen sollte,hängt vom persönlichen Urtyp ab: Nomade oder Ackerbauer?",
		}

		// Act
		p, err := m.FindByISBN("4545-8558-3232")
		if err != nil {
			return
		}

		// Assert
		assert.Equal(t, expectedP, p)
	})

}

func TestFindByTitle(t *testing.T) {

	books, _ := reader.LoadBooks("/mnt/d/projects/go-library/golang-kata-1/resources/books.csv")
	magazines, _ := reader.LoadMagazines("/mnt/d/projects/go-library/golang-kata-1/resources/magazines.csv")

	m := NewLibrarian(books, magazines)

	t.Run("CASE_SUCCESS_PRODUCT_FOUND", func(t *testing.T) {

		// Act
		res, err := m.FindByTitle("Ich helfe dir kochen")
		if err != nil {
			return
		}

		expectedResult := []models.Book{
			{
				ISBN: "5554-5545-4518",
			},
		}

		// Assert
		assert.Equal(t, expectedResult, res)
	})

	t.Run("CASE_FAIL_NONE_FOUND", func(t *testing.T) {

		// Act
		_, err := m.FindByTitle("RandomString1234556")
		if err != nil {
			return
		}

		// Assert
		assert.EqualError(t, liberror.ErrFailedToFindProduct, "failed to find product")
	})

	t.Run("CASE_FAIL_TITLE_NULL", func(t *testing.T) {

		// Act
		_, err := m.FindByTitle("")
		if err != nil {
			return
		}

		// Assert
		assert.EqualError(t, liberror.ErrFailedToFindProduct, "failed to find product")
	})

}

func TestFindBookByAuthor(t *testing.T) {

	// Arrange
	books, _ := reader.LoadBooks("/mnt/d/projects/go-library/golang-kata-1/resources/books.csv")
	magazines, _ := reader.LoadMagazines("/mnt/d/projects/go-library/golang-kata-1/resources/magazines.csv")

	m := NewLibrarian(books, magazines)

	t.Run("CASE_SUCCESS_FOUND_BOOKS", func(t *testing.T) {

		// Arrange
		expectedResult := []models.Book{
			{
				ISBN: "2365-8745-7854",
			},
		}

		// Act
		res, err := m.FindBookByAuthor("null-ferdinand@echocat.org")
		if err != nil {
			return
		}

		// Assert
		assert.Contains(t, expectedResult, res)

	})

	t.Run("CASE_FAIL_NO_BOOKS_FOUND", func(t *testing.T) {

		// Act
		_, err := m.FindBookByAuthor("")
		if err != nil {
			return
		}

		// Assert
		assert.EqualError(t, liberror.ErrInvalidEmail, "invalid email address")
	})

}

func TestFindMagazineByAuthor(t *testing.T) {

}
