package main

import (
	"fmt"
	service "gobooks/internal/services"
)

func main() {
	book := service.Book{
		ID: 1,
		Title: "The Hobbit",
		Author: "J.R.R. Tolkien",
		Genre: "Fantasy",
	}
	fmt.Println("The book is: ", book.Title)
}