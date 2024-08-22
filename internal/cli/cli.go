package cli

import service "gobooks/internal/services"

type BookCLI struct {
	service *service.BookService
}

func NewBookCLI(service *service.BookService) *BookCLI{
	return &BookCLI{service: service}
}