package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"git.vegaitsourcing.rs/radoslav.boychev/librarian-project/librarian"
	"git.vegaitsourcing.rs/radoslav.boychev/librarian-project/pkg/reader"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	start()
}

func start() {

	err := godotenv.Load("../.././config/config.env")
	if err != nil {
		log.Println(err)
		return
	}

	printMenu()

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

Loop:
	for {

		reader := bufio.NewReader(os.Stdin)

		char, _, err := reader.ReadRune()
		if err != nil {
			return
		}

		scanner := bufio.NewScanner(os.Stdin)

		switch char {
		case '1':
			fmt.Print("Insert a valid 14 digit ISBN divided by dashes '-' to search for: ")
			scanner.Scan()
			if err != nil {
				return
			}
			p, err := manager.FindByISBN(scanner.Text())
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			p.PrintProduct()
			fmt.Println()
			printMenu()

		case '2':
			fmt.Print("Title to search for:  ")
			scanner.Scan()
			if err != nil {
				return
			}
			_, err := manager.FindByTitle(scanner.Text())
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			fmt.Println()
			printMenu()

		case '3':
			fmt.Print("Author email to search a book for: ")
			scanner.Scan()
			if err != nil {
				return
			}
			_, err = manager.FindBookByAuthor(scanner.Text())
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			fmt.Println()
			printMenu()

		case '4':
			fmt.Print("Magazine author email to search for:  ")
			scanner.Scan()
			if err != nil {
				return
			}
			_, err = manager.FindMagazineByAuthor(scanner.Text())
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			fmt.Println()
			printMenu()

		case '5':
			fmt.Println("Exiting program")
			break Loop
		default:
			fmt.Println("Invalid choice. Please try again")
			fmt.Print("Your choice: ")
		}
	}
}

func printMenu() {
	fmt.Println("=-=-=-=-= MAIN MENU =-=-=-=-=")
	fmt.Println("Press '1' to find product by ISBN")
	fmt.Println("Press '2' to find by title")
	fmt.Println("Press '3' to find book by author")
	fmt.Println("Press '4' to find magazine by author")
	fmt.Println("Press '5' to exit")
	fmt.Print("Your choice: ")
}
