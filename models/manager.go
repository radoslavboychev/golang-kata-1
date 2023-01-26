package models

import (
	"strings"
)

type Manager struct {
	Books     []Book
	Magazines []Magazine
	Authors   []Author
}

// NewManager constructor creates a new product manager
func NewManager(books []Book, magazines []Magazine, authors []Author) *Manager {
	return &Manager{
		Books:     books,
		Magazines: magazines,
		Authors:   authors,
	}
}

func (m Manager) PrintBooks() {
	for _, bookIn := range m.Books {

		bookIn.PrintProduct()
	}
}

func (m Manager) PrintMagazines() error {
	for _, magIn := range m.Magazines {

		err := magIn.PrintProduct()
		if err != nil {
			return err
		}
	}
	return nil
}

func (m Manager) MapAuthors() {
	for _, book := range m.Books {
		for _, loadedAuthor := range m.Authors {
			for _, bookAuthor := range book.Authors {
				if strings.Contains(loadedAuthor.Email, bookAuthor.Email) {
					bookAuthor.Firstname = loadedAuthor.Firstname
					bookAuthor.Lastname = loadedAuthor.Lastname
					break
				}
			}
		}
	}
}

// FindByISBN looks up a product by ISBN and returns it
func (m *Manager) FindByISBN(isbn string) (Product, error) {
	var p Product

	for _, mag := range m.Magazines {
		if mag.ISBN == isbn {
			p = mag
		}
	}
	return p, nil
}

// // FindBookByAuthors should return all books where the author email matches
// func (m *Manager) FindByAuthor(email string) []Product {
// 	var res []Product

// 	// scan for books
// 	for _, b := range m.Books {
// 		if b.Authors == email {
// 			res = append(res, b)
// 			for _, auth := range m.Authors {
// 				if auth.Email == b.Authors {
// 					b.Authors = auth.Firstname + " " + auth.Lastname
// 				}
// 			}
// 			b.PrintProduct()
// 		}
// 	}

// 	// scan for magazines
// 	for _, mag := range m.Magazines {
// 		if mag.Authors == email {
// 			res = append(res, mag)
// 			for _, auth := range m.Authors {
// 				if auth.Email == mag.Authors {
// 					mag.Authors = auth.Firstname + " " + auth.Lastname
// 				}
// 			}
// 			mag.PrintProduct()
// 		}
// 	}
// 	return res
// }

// func (m Manager) PrintMagazines() {
// 	for _, magIn := range m.Magazines {
// 		for _, auth := range m.Authors {
// 			if auth.Email == magIn.Authors {
// 				magIn.Authors = auth.Firstname + " " + auth.Lastname
// 			}
// 		}
// 		magIn.PrintProduct()
// 	}
// }

func (m Manager) PrintAll() {
	m.PrintBooks()
	// m.PrintMagazines()
}
