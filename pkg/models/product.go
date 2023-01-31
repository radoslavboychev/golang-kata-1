package models

// Product interface defines behavior for products
type Product interface {
	PrintProduct() error
}

type Items struct {
	Magazines []Magazine
	Books     []Book
}
