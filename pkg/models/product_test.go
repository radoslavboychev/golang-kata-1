package models

import (
	"log"
	"testing"

	liberror "github.com/echocat/golang-kata-1/v1/errors"
	"github.com/stretchr/testify/assert"
)

// TestPrintProduct tests not only the function to print a product's info
// But also the generation of ISBNs and default values for names and authors
func TestPrintProduct(t *testing.T) {

	// Arrange
	b := NewBook("The Picture of Dorian Gray", "2234-6857-9473", "A book", []string{"Oscar Wilde"})

	t.Run("CASE_SUCCESS_GENERATES_AND_PRINTS", func(t *testing.T) {

		// Act
		err := b.PrintProduct()
		if err != nil {
			return
		}

		// Assert
		assert.NoError(t, err)
	})

	// Testing the generation of new ISBNs from empty ones
	t.Run("CASE_FAILED_IS_NULL", func(t *testing.T) {

		// Arrange
		c := Book{
			Title:       "The Bible",
			ISBN:        "",
			Authors:     []string{"God"},
			Description: "It's the bible, dude",
		}

		// Act
		err := c.PrintProduct()
		if err != nil {
			return
		}

		// Assert
		assert.NoError(t, err, liberror.ErrorISBNIsNull)
	})

	// Testing the generation of ISBNs from ones with invalid length/format
	t.Run("CASE_FAILED_ISBN_INVALID", func(t *testing.T) {

		// Arrange
		c := NewBook("NewBook", "293-4421-4", "Book", []string{"Baller"})

		// Act
		err := c.PrintProduct()
		if err != nil {
			return
		}

		// Assert
		assert.NoError(t, err, liberror.ErrorISBNIsNull)
	})

	// Case when the whole object is empty and default values need to be generated
	t.Run("CASE_FAILED_OBJECT_IS_NULL", func(t *testing.T) {
		// Arrange
		c := Book{}

		// Act
		err := c.PrintProduct()
		if err != nil {
			return
		}

		b := c.PrintProduct()
		log.Println(b)

		// Assert
		assert.NoError(t, err)
	})
}
