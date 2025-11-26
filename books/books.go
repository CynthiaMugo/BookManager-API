package books

import (
	"fmt"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func AddBook(list []Book, title, author string) []Book {
	newBook := Book{
		ID:     len(list) + 1,
		Title:  title,
		Author: author,
	}
	list = append(list, newBook)
	fmt.Println("Book added:", newBook.Title)
	return list
}

func UpdateBook(list []Book, id int, title, author string) ([]Book, bool) {
	for i, b := range list {
		if b.ID == id {
			list[i].Title = title
			list[i].Author = author
			return list, true
		}
	}
	return list, false
}

func DeleteBook(list []Book, id int) ([]Book, bool) {
	for i, b := range list {
		if b.ID == id {
			list = append(list[:i], list[i+1:]...)
			return list, true
		}
	}
	return list, false
}
