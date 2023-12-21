package Search

import (
	"AAtries/Book"
	"fmt"
)

func SearchByTitle(title string) {
	fmt.Println(Book.GetBooks())
	fmt.Printf("Search Results for Title '%s':\n", title)
	for _, book := range Book.GetBooks() {
		if book.GetTitle() == title {
			fmt.Printf("Author: %s, Category: %s\n", book.Author, book.Category)
		}
	}
}

/*
	func SearchByTitle(title string) []Book.Book {
		var matchedBooks []Book.Book
		for _, book := range Book.GetBooks() {
			if book.GetTitle() == title {
				matchedBooks = append(matchedBooks, book)
			}
		}
		return matchedBooks
	}
*/
func SearchByAuthor(author string) {
	fmt.Println(Book.GetBooks())
	fmt.Printf("Search Results for Author '%s':\n", author)
	for _, book := range Book.GetBooks() {
		if book.GetAuthor() == author {
			fmt.Printf("Title: %s, Category: %s\n", book.Title, book.Category)
		}
	}
}

func SearchByCategory(category string) {
	fmt.Println(Book.GetBooks())
	fmt.Printf("Search Results for Category '%s':\n", category)
	for _, book := range Book.GetBooks() {
		if book.GetCategory() == category {
			fmt.Printf("Title: %s, Author: %s\n", book.Title, book.Author)
		}
	}
}
