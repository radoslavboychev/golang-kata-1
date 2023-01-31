package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	liberror "github.com/echocat/golang-kata-1/v1/errors"
	"github.com/echocat/golang-kata-1/v1/librarian"
	"github.com/echocat/golang-kata-1/v1/pkg/reader"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	liberror "github.com/radoslav.boychev/librarian-projecterrors"
	"github.com/radoslav.boychev/librarian-projectlibrarian"
	"github.com/radoslav.boychev/librarian-projectpkg/reader"
)

func main() {

	// initialize the app
	start()
}

func start() {

	err := godotenv.Load("../.././config/config.env")
	if err != nil {
		log.Println(err)
		return
	}

	// ENV
	magazinesFile := os.Getenv("MAGAZINES_FILE")
	booksFile := os.Getenv("BOOKS_FILE")

	// Load books from file
	books, err := reader.LoadBooks(magazinesFile)
	if err != nil {
		log.Println(err)
		return
	}

	// Load magazines from file
	magazines, err := reader.LoadMagazines(booksFile)
	if err != nil {
		log.Println(err)
		return
	}

	// creates a new products manager (librarian)
	manager := librarian.NewLibrarian(books, magazines)

	// reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)

	var choice int

	for choice < 6 {

		// print the menu for reaction
		printMenu()

		fmt.Scanln(&choice)
		if err != nil {
			return
		}

		switch choice {
		case 1:
			fmt.Print("Insert a valid 14 digit ISBN divided by dashes '-' to search for: ")
			scanner.Scan()
			if err != nil {
				return
			}
			p, err := manager.FindByISBN(scanner.Text())
			if err != nil {
				if err == liberror.ErrFailedToFindProduct {
					fmt.Println(err.Error())
					break
				} else {
					log.Fatalf(err.Error())
					break
				}
			}

			p.PrintProduct()
			fmt.Println()

		case 2:
			fmt.Print("Title to search for:  ")
			scanner.Scan()
			if err != nil {
				return
			}
			_, err := manager.FindByTitle(scanner.Text())
			if err != nil {
				if err == liberror.ErrFailedToFindProduct {
					fmt.Println(err.Error())
					break
				}
			}
			fmt.Println()

		case 3:
			fmt.Print("Author email to search a book for: ")
			scanner.Scan()
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			_, err = manager.FindBookByAuthor(scanner.Text())
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			fmt.Println()

		case 4:
			fmt.Print("Magazine author email to search for:  ")
			scanner.Scan()
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			_, err = manager.FindMagazineByAuthor(scanner.Text())
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			fmt.Println()

		case 5:
			fmt.Println("Sorting items...")
			_, err := manager.Sort()
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			fmt.Println()

		case 6:
			fmt.Println("Exiting program...")
			os.Exit(0)

		default:
			fmt.Println("Invalid input. Exiting program...")
			os.Exit(0)
		}

	}
}

// print menu
func printMenu() {
	fmt.Println("=-=-=-=-= MAIN MENU =-=-=-=-=")
	fmt.Println("Press '1' to find product by ISBN")
	fmt.Println("Press '2' to find by title")
	fmt.Println("Press '3' to find book by author")
	fmt.Println("Press '4' to find magazine by author")
	fmt.Println("Press '5' to sort items by title")
	fmt.Println("Press '6' to exit")
	fmt.Print("Your choice: ")
}
