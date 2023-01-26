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

	// // LOAD AUTHOR RECORDS
	// authors, err := LoadAuthors("./resources/authors.csv")
	// if err != nil {
	// 	return nil, err
	// }

	// out := []models.Book{}

	// SLICE TO ITERATE
	res := []models.Book{}

	// RANGE THROUGH RECORDS SLICE
	// CREATE A BOOK FOR EVERY ENTRY
	for _, rec := range records {
		b := models.Book{
			Title: rec[0],
			ISBN:  rec[1],
			Authors: []models.Author{
				{
					Email: rec[2],
				},
			},
			Description: rec[3],
		}

		res = append(res, b)
	}

	// for _, b := range res {
	// 	for _, bookAuthor := range b.Authors {
	// 		for _, loadedAuthor := range authors {
	// 			if bookAuthor.Email == loadedAuthor.Email {
	// 				bookAuthor.Firstname = loadedAuthor.Firstname
	// 				bookAuthor.Lastname = loadedAuthor.Lastname
	// 				out = append(out, b)
	// 			}
	// 		}

	// 	}
	// }

	return res, nil
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

	// SLICE TO ITERATE
	res := []models.Magazine{}

	// RANGE THROUGH RECORDS SLICE
	// CREATE A BOOK FOR EVERY ENTRY
	for _, rec := range records {
		m := models.Magazine{
			Title:       rec[0],
			ISBN:        rec[1],
			Authors:     []string{rec[2]},
			PublishedAt: rec[3],
		}

		res = append(res, m)
	}

	return res, nil

}

// ResolveAuthors
func ResolveAuthors(authors []models.Author, magazines []models.Magazine) []models.Magazine {

	var mappedMagazines []models.Magazine
	for _, mag := range magazines {
		for _, a := range mag.Authors {
			newAuthors := splitString(a)
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
	return mappedMagazines
}

// LoadAuthors reads all authors
func LoadAuthors(filename string) ([]models.Author, error) {
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

func splitString(line string) []string {
	res := strings.Split(line, ",")
	return res
}
