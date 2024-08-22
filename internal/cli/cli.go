package cli

import (
	"fmt"
	"os"
	"strconv"
	"time"

	service "gobooks/internal/services"
)

type BookCLI struct {
	service *service.BookService
}

func NewBookCLI(service *service.BookService) *BookCLI{
	return &BookCLI{service: service}
}

func (cli *BookCLI) Run(){
	if len(os.Args) < 2{
		fmt.Println("Usage: books <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books search <book title>")
			return
		}
		bookName := os.Args[2]
		cli.searchBooks(bookName)
	case "simulate":
		if len(os.Args) < 3 {
			fmt.Println("Usage: books simulate <book_id> <book_id> <book_id> ...")
			return
		}
		bookIDs := os.Args[2:]
		cli.simulateReading(bookIDs)

	}

	
}

func (cli *BookCLI) searchBooks(name string){
	books, err := cli.service.SearchBooksByName(name)
	if err != nil {
		fmt.Println("Error searching books:", err )
		return
	}
	if len(books) == 0 {
		fmt.Println("No books found.")
		return
	}

	fmt.Printf("%d books found\n ", len(books))

	for _, book := range books{
		fmt.Printf("ID: %d, Title: %s, Author: %s, Genre: %s\n",
			book.ID, book.Title, book.Author, book.Genre,
		)
	}
}

func (cli *BookCLI) simulateReading(booksIDsSrt []string){
	var bookIDs []int
	for _, idStr := range booksIDsSrt{
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid book ID: ", idStr)
			continue
		}
		bookIDs = append(bookIDs, id)
	}
	responses := cli.service.SimulateMultipleReadings(bookIDs, 2*time.Second)

	for _, response := range responses {
		fmt.Println(response)
	}
}