package reader

import (
	"log"
	"os"
	"testing"

	liberror "github.com/echocat/golang-kata-1/v1/errors"
	"github.com/echocat/golang-kata-1/v1/pkg/models"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// ENV
var _ = godotenv.Load("../.././config/config.env")
var magazinesFile = os.Getenv("MAGAZINES_FILE")
var magazinesTest = os.Getenv("MAGAZINES_TEST_FILE")
var booksFile = os.Getenv("BOOKS_FILE")
var booksTest = os.Getenv("BOOKS_TEST")
var authorsFile = os.Getenv("AUTHORS_FILE")
var authorsTestFile = os.Getenv("AUTHORS_TEST_FILE")

// TestLoadBooks is testing the LoadBooks function to lead books from files
func TestLoadBooks(t *testing.T) {

	// Case when books file can not be found
	t.Run("CASE_FAILED_FILE_NOT_FOUND", func(t *testing.T) {
		bookPath := "./file.csv"
		_, err := LoadBooks(bookPath)
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrFailedToOpenFile)
	})

	// Case when the file name provided is empty
	t.Run("CASE_FAILED_FILENAME_IS_NULL", func(t *testing.T) {
		bookPath := ""
		_, err := LoadBooks(bookPath)
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrFilenameInvalid)
	})

	// Case when no books have been found
	t.Run("CASE_FAILED_NO_BOOKS_FOUND", func(t *testing.T) {
		_, err := LoadBooks(booksTest)
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrNoBooksLoaded)

	})

	// Case when books have been loaded
	t.Run("CASE_SUCCESS_BOOKS_LOADED", func(t *testing.T) {
		_, err := LoadBooks(booksFile)
		if err != nil {
			return
		}

		assert.NoError(t, err)
	})
}

// TestLoadMagazines loads the magazines from a file
func TestLoadMagazines(t *testing.T) {

	// Case when the file name is empty
	t.Run("CASE_FAILED_FILENAME_NULL", func(t *testing.T) {
		filename := ""

		_, err := LoadMagazines(filename)
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrFilenameInvalid)

	})

	// Case when file is not found
	t.Run("CASE_FAILED_FILE_NOT_FOUND", func(t *testing.T) {
		_, err := LoadMagazines("./file.csv")
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrFileNotFound)
	})

	// Case when no magazines are found
	t.Run("CASE_FAILED_MAGAZINES_NOT_FOUND", func(t *testing.T) {
		_, err := LoadMagazines(magazinesTest)
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrNoMagazinesLoaded)
	})

	// Case when a magazine is successfully found
	t.Run("CASE_SUCCESS_MAGAZINES_FOUND", func(t *testing.T) {
		_, err := LoadMagazines(magazinesFile)
		if err != nil {
			return
		}

		assert.NoError(t, err)
	})
}

// TestResolveMagAuthors testing resolving magazine authors to names
func TestResolveMagAuthors(t *testing.T) {

	// Arrange
	magazines, err := LoadMagazines(booksFile)
	if err != nil {
		return
	}

	authors, err := LoadAuthors(authorsFile)
	if err != nil {
		return
	}

	// Case when the authors provided are a null struct
	t.Run("CASE_FAILED_AUTHORS_NULL", func(t *testing.T) {

		authors := []models.Author{}

		// Act
		_, err := ResolveMagAuthors(authors, magazines)
		if err != nil {
			return
		}

		// Assert
		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrFailedToResolveAuthorsInvalid)

	})

	// Case when the magazines are null
	t.Run("CASE_FAILED_MAGAZINES_NULL", func(t *testing.T) {

		magazines := []models.Magazine{}

		// Act
		_, err := ResolveMagAuthors(authors, magazines)
		if err != nil {
			return
		}

		// Assert
		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrFailedToResolveMagazinesInvalid)
	})

	// Case when authors are successfully mapped
	t.Run("CASE_SUCCESS_AUTHORS_MAPPED", func(t *testing.T) {
		_, err := ResolveMagAuthors(authors, magazines)
		if err != nil {
			return
		}

		assert.NoError(t, err)

	})
}

// TestResolveBookAuthors tests
func TestResolveBookAuthors(t *testing.T) {

	// Load books from file
	books, err := LoadBooks(magazinesFile)
	if err != nil {
		log.Println(err)
		return
	}

	// Load authors from file
	authors, err := LoadAuthors(authorsFile)
	if err != nil {
		return
	}

	// Case when the authors struct provided is null
	t.Run("CASE_FAILED_AUTHORS_NULL", func(t *testing.T) {

		authors = []models.Author{}

		_, err := ResolveBookAuthors(authors, books)
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrFailedToResolveMagazinesInvalid)

	})

	// Case when books provided are a null struct
	t.Run("CASE_FAILED_BOOKS_NULL", func(t *testing.T) {

		books = []models.Book{}

		_, err := ResolveBookAuthors(authors, books)
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrFailedToResolveMagazinesInvalid)

	})

	// Case when authors are successfully mapped
	t.Run("CASE_SUCCESS_AUTHORS_MAPPED", func(t *testing.T) {
		_, err := ResolveBookAuthors(authors, books)
		if err != nil {
			return
		}

		assert.NoError(t, err)
	})
}

// TestLoadAuthors tests loading author data from file
func TestLoadAuthors(t *testing.T) {

	// Case when file name provided is invalid
	t.Run("CASE_FAILED_FILENAME_INVALID", func(t *testing.T) {
		_, err := LoadAuthors("")
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrFilenameInvalid)
	})

	// Case when the file provided is empty
	t.Run("CASE_FAILED_FILE_EMPTY", func(t *testing.T) {

		_, err := LoadAuthors(authorsTestFile)
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrFailedToOpenFile)
	})

	// Case when file is not found
	t.Run("CASE_FAILED_FILE_NOT_FOUND", func(t *testing.T) {
		_, err := LoadAuthors(".")
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrFailedToOpenFile)
	})

	// Case when authors are successfully loaded
	t.Run("CASE_SUCCESS_AUTHORS_LOADED", func(t *testing.T) {
		_, err := LoadAuthors(authorsFile)
		if err != nil {
			return
		}

		assert.NoError(t, err)

	})
}
