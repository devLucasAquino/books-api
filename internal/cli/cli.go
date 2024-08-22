package cli

import (
	"fmt"
	"os"

	service "gobooks/internal/services"
)

type BookCLI struct {
	service *service.BookService
}

func NewBookCLI(service *service.BookService) *BookCLI{
	return &BookCLI{service: service}
}

func (cli *BookCLI) run(){
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
			book.ID, book.Title, book.Author, book.Genre
		)
	}
}