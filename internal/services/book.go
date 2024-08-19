package service

type Book struct {
	ID int
	Title string
	Author string
	Genre string
}

func (b Book) GetFullBook() string {
	return b.Title + " by " + b.Author
}