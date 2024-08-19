package main

import (
	"database/sql"
	service "gobooks/internal/services"
	"gobooks/internal/web"
	"net/http"
)

func main() {

	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	bookService := service.NewBookService(db)
	BookHandlers := web.NewBookHandlers(bookService)

	router := http.NewServeMux()
	router.HandleFunc("GET /books", BookHandlers.GetBooks)
	router.HandleFunc("POST /books", BookHandlers.CreateBook)
	router.HandleFunc("GET /books/{id}", BookHandlers.GetBookByID)
	router.HandleFunc("PUT /books/{id}", BookHandlers.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", BookHandlers.DeleteBook)

	http.ListenAndServe(":8080", router)
}