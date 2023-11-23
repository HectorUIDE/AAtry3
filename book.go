package main

import "fmt"

type Book struct {
	Title    string
	Author   string
	Category string
}

var Books []Book

func AddBook() {
	var newBook Book
	fmt.Print("Enter Title: ")
	fmt.Scanln(&newBook.Title)
	fmt.Print("Enter Author: ")
	fmt.Scanln(&newBook.Author)
	fmt.Print("Enter Category: ")
	fmt.Scanln(&newBook.Category)

	Books = append(Books, newBook)
	fmt.Println("Book added successfully!")
}

func ViewBooks() {
	fmt.Println("Books:")
	for _, book := range Books {
		fmt.Printf("Title: %s, Author: %s, Category: %s\n", book.Title, book.Author, book.Category)
	}
}
