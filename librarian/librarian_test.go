package librarian

import (
	"os"
	"testing"

	liberror "git.vegaitsourcing.rs/radoslav.boychev/librarian-project/errors"
	"git.vegaitsourcing.rs/radoslav.boychev/librarian-project/pkg/models"
	"git.vegaitsourcing.rs/radoslav.boychev/librarian-project/pkg/reader"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var err = godotenv.Load(".././config/config.env")

// ENV
var magazinesFile = os.Getenv("MAGAZINES_FILE")
var booksFile = os.Getenv("BOOKS_FILE")

// Testing looking up a product by ISBN
func TestFindByISBN(t *testing.T) {

	// Arrange
	books, _ := reader.LoadBooks(booksFile)
	magazines, _ := reader.LoadMagazines(magazinesFile)
	m := NewLibrarian(books, magazines)

	// Case when the ISBN number is invalid length
	t.Run("CASE_ISBN_INVALID", func(t *testing.T) {

		// Act
		_, err := m.FindByISBN("4545-8558-3232-22948")
		if err != nil {
			return
		}

		// Assert
		assert.EqualError(t, liberror.ErrISBNInvalid, "ISBN is of invalid length")

	})

	// Case when the ISBN is valid but no results have been found
	t.Run("CASE_ISBN_VALID_NO_RESULTS", func(t *testing.T) {

		// Act
		m.FindByISBN("1024-1024-1024")

		// Assert
		assert.EqualError(t, liberror.ErrFailedToFindProduct, "failed to find product")

	})

	// Case when the ISBN is valid and a product is found
	t.Run("CASE_ISBN_VALID_RESULT_FOUND", func(t *testing.T) {

		expectedP := models.Book{
			Title:       "Schlank im Schlaf ",
			ISBN:        "4545-8558-3232",
			Authors:     []string{" Karl Gustafsson "},
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

// Testing looking up a product by title
func TestFindByTitle(t *testing.T) {

	// Arrange
	books, _ := reader.LoadBooks(booksFile)
	magazines, _ := reader.LoadMagazines(magazinesFile)

	m := NewLibrarian(books, magazines)

	// Case when a product has been found by title
	t.Run("CASE_SUCCESS_PRODUCT_FOUND", func(t *testing.T) {

		// Act
		res, err := m.FindByTitle("Ich helfe dir kochen")
		if err != nil {
			return
		}

		expectedResult := []models.Product([]models.Product{models.Book{Title: "Ich helfe dir kochen. Das erfolgreiche Universalkochbuch mit großem Backteil", ISBN: "5554-5545-4518", Authors: []string{" Paul Walter "}, Description: "Auf der Suche nach einem Basiskochbuch steht man heutzutage vor einer Fülle von Alternativen. Es fällt schwer,daraus die für sich passende Mixtur aus Grundlagenwerk und Rezeptesammlung zu finden. Man sollte sich darüber im Klaren sein,welchen Schwerpunkt man setzen möchte oder von welchen Koch- und Backkenntnissen man bereits ausgehen kann."}})

		// Assert
		assert.Equal(t, expectedResult, res)
	})

	// Case when no products have been found
	t.Run("CASE_FAIL_NONE_FOUND", func(t *testing.T) {

		// Act
		_, err := m.FindByTitle("RandomString1234556")
		if err != nil {
			return
		}

		// Assert
		assert.ErrorIs(t, liberror.ErrFailedToFindProduct, err)
	})

	// Case when the title is null
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

// Testing finding a book by its author
func TestFindBookByAuthor(t *testing.T) {

	// Arrange
	books, _ := reader.LoadBooks(booksFile)
	magazines, _ := reader.LoadMagazines(magazinesFile)

	m := NewLibrarian(books, magazines)

	// Case when a book has been found
	t.Run("CASE_SUCCESS_FOUND_BOOKS", func(t *testing.T) {

		// Act
		_, err := m.FindBookByAuthor("null-ferdinand@echocat.org")
		if err != nil {
			return
		}

		// Assert
		assert.NoError(t, err)

	})

	// Case when no books have been found
	t.Run("CASE_FAIL_NO_BOOKS_FOUND", func(t *testing.T) {

		// Act
		_, err := m.FindBookByAuthor("")
		if err != nil {
			return
		}

		// Assert
		assert.Error(t, err, liberror.ErrInvalidEmail)
		assert.EqualError(t, liberror.ErrInvalidEmail, "invalid email address")
	})

}

// Testing looking up a magazine by author
func TestFindMagazineByAuthor(t *testing.T) {

	// Arrange
	books, _ := reader.LoadBooks(booksFile)
	magazines, _ := reader.LoadMagazines(magazinesFile)
	m := NewLibrarian(books, magazines)

	// Case when email is empty
	t.Run("CASE_FAILS_EMAIL_IS_NIL", func(t *testing.T) {

		// Act
		_, err := m.FindMagazineByAuthor("")
		if err != nil {
			return
		}

		// Assert
		assert.Error(t, err, liberror.ErrFailedToPrint)
		assert.ErrorIs(t, err, liberror.ErrEmailIsNull)
	})

	// Case when no product is found
	t.Run("CASE_FAILS_NO_PRODUCT_FOUND", func(t *testing.T) {

		// Act
		_, err := m.FindMagazineByAuthor("example@gmail.com")
		if err != nil {
			return
		}

		// Assert
		assert.Error(t, err)
		assert.ErrorIs(t, liberror.ErrEmailIsNull, err)

	})

	// Case when a product is successfully found
	t.Run("CASE_SUCCESS_PRODUCT_FOUND", func(t *testing.T) {

		// Act
		_, err := m.FindMagazineByAuthor("null-mueller@echocat.org")
		if err != nil {
			return
		}

		// Assert
		assert.NoError(t, err)

	})

}
