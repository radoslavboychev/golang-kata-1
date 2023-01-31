package reader

import (
	"log"
	"os"
	"testing"

	liberror "git.vegaitsourcing.rs/radoslav.boychev/librarian-project/errors"
	"git.vegaitsourcing.rs/radoslav.boychev/librarian-project/pkg/models"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

<<<<<<< HEAD
=======
var _ = godotenv.Load("../.././config/config.env")

>>>>>>> dbf35c6 (console menu, test cases)
// ENV
var _ = godotenv.Load("../.././config/config.env")
var magazinesFile = os.Getenv("MAGAZINES_FILE")
var magazinesTest = os.Getenv("MAGAZINES_TEST_FILE")
var booksFile = os.Getenv("BOOKS_FILE")
var booksTest = os.Getenv("BOOKS_TEST")
var authorsFile = os.Getenv("AUTHORS_FILE")
var authorsTestFile = os.Getenv("AUTHORS_TEST_FILE")

<<<<<<< HEAD
// TestLoadBooks is testing the LoadBooks function to lead books from files
func TestLoadBooks(t *testing.T) {

	// Case when books file can not be found
=======
// Testing loading books from a file
func TestLoadBooks(t *testing.T) {

	// Case when a file to load is not found
>>>>>>> dbf35c6 (console menu, test cases)
	t.Run("CASE_FAILED_FILE_NOT_FOUND", func(t *testing.T) {
		bookPath := "./file.csv"
		_, err := LoadBooks(bookPath)
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrFailedToOpenFile)
	})

<<<<<<< HEAD
	// Case when the file name provided is empty
=======
	// Case when file name is null
>>>>>>> dbf35c6 (console menu, test cases)
	t.Run("CASE_FAILED_FILENAME_IS_NULL", func(t *testing.T) {
		bookPath := ""
		_, err := LoadBooks(bookPath)
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrFilenameInvalid)
	})

<<<<<<< HEAD
	// Case when no books have been found
=======
	// Case when no books are found in the file
>>>>>>> dbf35c6 (console menu, test cases)
	t.Run("CASE_FAILED_NO_BOOKS_FOUND", func(t *testing.T) {
		_, err := LoadBooks(booksTest)
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrNoBooksLoaded)

	})

<<<<<<< HEAD
	// Case when books have been loaded
=======
	// Case when all books are successfully loaded
>>>>>>> dbf35c6 (console menu, test cases)
	t.Run("CASE_SUCCESS_BOOKS_LOADED", func(t *testing.T) {
		_, err := LoadBooks(booksFile)
		if err != nil {
			return
		}

		assert.NoError(t, err)
	})
}

<<<<<<< HEAD
// TestLoadMagazines loads the magazines from a file
func TestLoadMagazines(t *testing.T) {

	// Case when the file name is empty
=======
// Testing loading magazines from a file
func TestLoadMagazines(t *testing.T) {

	// Case when filename is null
>>>>>>> dbf35c6 (console menu, test cases)
	t.Run("CASE_FAILED_FILENAME_NULL", func(t *testing.T) {
		filename := ""

		_, err := LoadMagazines(filename)
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrFilenameInvalid)

	})

<<<<<<< HEAD
	// Case when file is not found
=======
	// Case when the file is not found
>>>>>>> dbf35c6 (console menu, test cases)
	t.Run("CASE_FAILED_FILE_NOT_FOUND", func(t *testing.T) {
		_, err := LoadMagazines("./file.csv")
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrFileNotFound)
	})

<<<<<<< HEAD
	// Case when no magazines are found
=======
	// Case when the the file is empty
>>>>>>> dbf35c6 (console menu, test cases)
	t.Run("CASE_FAILED_MAGAZINES_NOT_FOUND", func(t *testing.T) {
		_, err := LoadMagazines(magazinesTest)
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrNoMagazinesLoaded)
	})

<<<<<<< HEAD
	// Case when a magazine is successfully found
=======
	// Case when magazines have been loaded successfully
>>>>>>> dbf35c6 (console menu, test cases)
	t.Run("CASE_SUCCESS_MAGAZINES_FOUND", func(t *testing.T) {
		_, err := LoadMagazines(magazinesFile)
		if err != nil {
			return
		}

		assert.NoError(t, err)
	})
}

<<<<<<< HEAD
// TestResolveMagAuthors testing resolving magazine authors to names
=======
// Testing resolving authors for magazines
>>>>>>> dbf35c6 (console menu, test cases)
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

<<<<<<< HEAD
	// Case when the authors provided are a null struct
=======
	// Case when authors struct is null
>>>>>>> dbf35c6 (console menu, test cases)
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

<<<<<<< HEAD
	// Case when the magazines are null
=======
	// Case when magazines struct is null
>>>>>>> dbf35c6 (console menu, test cases)
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

<<<<<<< HEAD
	// Case when authors are successfully mapped
=======
	// Case when authors have successfully been mapped
>>>>>>> dbf35c6 (console menu, test cases)
	t.Run("CASE_SUCCESS_AUTHORS_MAPPED", func(t *testing.T) {
		_, err := ResolveMagAuthors(authors, magazines)
		if err != nil {
			return
		}

		assert.NoError(t, err)

	})
}

<<<<<<< HEAD
// TestResolveBookAuthors tests
=======
// Testing resolving authors for books
>>>>>>> dbf35c6 (console menu, test cases)
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

<<<<<<< HEAD
	// Case when the authors struct provided is null
=======
	// Case when authors struct is null
>>>>>>> dbf35c6 (console menu, test cases)
	t.Run("CASE_FAILED_AUTHORS_NULL", func(t *testing.T) {

		authors = []models.Author{}

		_, err := ResolveBookAuthors(authors, books)
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrFailedToResolveMagazinesInvalid)

	})

<<<<<<< HEAD
	// Case when books provided are a null struct
=======
	// Case when books struct is null
>>>>>>> dbf35c6 (console menu, test cases)
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

<<<<<<< HEAD
// TestLoadAuthors tests loading author data from file
=======
// Testing loading authors from files
>>>>>>> dbf35c6 (console menu, test cases)
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

<<<<<<< HEAD
	// Case when the file provided is empty
=======
	// Case when the provided file is empty
>>>>>>> dbf35c6 (console menu, test cases)
	t.Run("CASE_FAILED_FILE_EMPTY", func(t *testing.T) {

		_, err := LoadAuthors(authorsTestFile)
		if err != nil {
			return
		}

		assert.Error(t, err)
		assert.ErrorIs(t, err, liberror.ErrFailedToOpenFile)
	})

<<<<<<< HEAD
	// Case when file is not found
=======
	// Case when no file has been found
>>>>>>> dbf35c6 (console menu, test cases)
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
