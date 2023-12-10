package Book

import (
	"fmt"
)

type Book struct {
	Title    string
	Author   string
	Category string
}

var books []Book

//--------------Setters------------------------------------

// Title setter, it also controls the title's length
func (b *Book) SetTitle(title string) error {
	for _, book := range books {
		if book.GetTitle() == title {
			return fmt.Errorf("Title already exists")
		}
	}
	b.Title = title
	return fmt.Errorf("Title is adequate!")
}

// Author setter, it also controls the author's length
func (b *Book) SetAuthor(author string) error {
	if len(author) > 10 {
		return fmt.Errorf("Author too long")
	}
	b.Author = author
	return fmt.Errorf("Author is adequate!")
}

// Cuthor setter, it also controls the author's length
func (b *Book) SetCategory(category string) error {
	if len(category) > 10 {
		return fmt.Errorf("Category too long")
	}
	b.Category = category
	return fmt.Errorf("Category is adequate!")
}

//--------------Getters------------------------------------

func (b *Book) GetTitle() string {
	return b.Title
}

func (b *Book) GetAuthor() string {
	return b.Author
}

func (b *Book) GetCategory() string {
	return b.Category
}

//-----------Functions-------------------------------------------------------------

// This function returns the value of the slice books (*it doesn't modify it, it only returns a copy)
func GetBooks() []Book {
	return books
}

// This special function allows me to modify the value itself of the slice book
func GetBooksPointer() *[]Book {
	return &books
}

func (newBook Book) AddBook() error {
	if newBook.GetTitle() == "" || newBook.GetAuthor() == "" || newBook.GetCategory() == "" {
		return fmt.Errorf("Book failed to be added")
	}
	fmt.Println(newBook)
	books = append(books, newBook)
	return fmt.Errorf("Book added successfully!")
}

func ViewBooks() {
	fmt.Println("Books:")
	fmt.Println(books)
	for _, book := range books {
		fmt.Printf("Title: %s, Author: %s, Category: %s\n", book.Title, book.Author, book.Category)
	}
}
