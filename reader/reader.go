package reader

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/echocat/golang-kata-1/v1/models"
)

// LoadBooks reads book data from a file
func LoadBooks(filename string) ([]models.Book, error) {

	// OPEN FILE
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

	a, err := loadAuthors("./resources/authors.csv")
	if err != nil {
		return nil, err
	}

	out, err := resolveBookAuthors(a, res)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// // LoadMagazines loads magazine data from a file
func LoadMagazines(filename string) ([]models.Magazine, error) {
	// read magazine from file
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

	res := []models.Magazine{}
	for _, rec := range records {
		m := models.Magazine{
			Title:       rec[0],
			ISBN:        rec[1],
			Authors:     []string{rec[2]},
			PublishedAt: rec[3],
		}

		res = append(res, m)
	}

	a, err := loadAuthors("./resources/authors.csv")
	if err != nil {
		return nil, err
	}

	out, err := resolveMagAuthors(a, res)
	if err != nil {
		return nil, err
	}

	return out, nil

}

// resolveMagAuthirs maps authors to their emails for magazines
func resolveMagAuthors(authors []models.Author, magazines []models.Magazine) (mag []models.Magazine, err error) {

	var mappedMagazines []models.Magazine
	for _, mag := range magazines {
		for _, a := range mag.Authors {
			newAuthors := splitString(a, ",")
			mag.Authors = []string{}
			for _, mappedAuthor := range newAuthors {
				for _, auth := range authors {
					if mappedAuthor == auth.Email {
						mappedAuthor = " " + auth.Firstname + " " + auth.Lastname + " "

						mag.Authors = append(mag.Authors, mappedAuthor)

					}

				}
			}

		}
		mappedMagazines = append(mappedMagazines, mag)
	}
	return mappedMagazines, nil
}

// resolveBookAuthors maps authors email to their name
func resolveBookAuthors(authors []models.Author, magazines []models.Book) (mag []models.Book, err error) {
	var mappedBooks []models.Book
	for _, mag := range magazines {
		for _, a := range mag.Authors {
			newAuthors := splitString(a, ",")
			mag.Authors = []string{}
			for _, mappedAuthor := range newAuthors {
				for _, auth := range authors {
					if mappedAuthor == auth.Email {
						mappedAuthor = " " + auth.Firstname + " " + auth.Lastname + " "

						mag.Authors = append(mag.Authors, mappedAuthor)

					}

				}
			}

		}
		mappedBooks = append(mappedBooks, mag)
	}
	return mappedBooks, nil
}

// LoadAuthors reads all authors
func loadAuthors(filename string) ([]models.Author, error) {
	// reads the authors from file
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
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

	// maps them to struct
	var authors []models.Author

	for _, author := range records {

		a := models.Author{
			Email:     author[0],
			Firstname: author[1],
			Lastname:  author[2],
		}
		authors = append(authors, a)
	}
	return authors, nil
}

// splitStrings separates a line of strings into a slice of strings
func splitString(line, symbol string) []string {
	res := strings.Split(line, symbol)
	return res
}
