package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/CynthiaMugo/BookManager-API/books"
)

const dataFile = "books.json"

func LoadBooks() []books.Book {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		return []books.Book{}
	}

	var list []books.Book
	if err := json.Unmarshal(file, &list); err != nil {
		fmt.Println("Error reading JSON:", err)
		return []books.Book{}
	}
	return list
}

func SaveBooks(list []books.Book) {
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	if err := os.WriteFile(dataFile, data, 0644); err != nil {
		fmt.Println("Error writing JSON:", err)
	}
}
