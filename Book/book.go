package Book

import (
	"encoding/json"
	"fmt"
)

type JsonBook struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

type Book struct {
	Title    string
	Author   string
	Category string
}

var books []Book

//--------------Setters------------------------------------

// Title setter, it makes sure that no titles are repeated
func (b *Book) SetTitle(title string) error {
	for _, book := range books {
		if book.GetTitle() == title {
			return fmt.Errorf("Title already exists")
		}
	}
	b.Title = title
	fmt.Println("Title is adequate!")
	return nil
}

// Author setter, it also controls the author's length
func (b *Book) SetAuthor(author string) error {
	if len(author) > 10 {
		return fmt.Errorf("Author too long")
	}
	b.Author = author
	fmt.Println("Author is adequate!")
	return nil
}

// Cuthor setter, it also controls the author's length
func (b *Book) SetCategory(category string) error {
	if len(category) > 10 {
		return fmt.Errorf("Category too long")
	}
	b.Category = category
	fmt.Println("Category is adequate!")
	return nil
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

// This function only adds a new book to the collection as long as it checks
// all the requirements (Title, Author, and Category)
func (newBook Book) AddBook() error {
	if newBook.GetTitle() == "" || newBook.GetAuthor() == "" || newBook.GetCategory() == "" {
		return fmt.Errorf("Book failed to be added")
	}
	fmt.Println(newBook)
	books = append(books, newBook)

	// Serialize result and print it to console
	JsonSerialize(newBook)

	//Finish
	return fmt.Errorf("Book added successfully!")
}

func ViewBooks() {
	fmt.Println("Books:")
	//	fmt.Println(books)
	for i, book := range books {
		// Serialize result and print it to console
		fmt.Printf("%d. ", i)
		JsonSerialize(book)
		//		fmt.Printf("Title: %s, Author: %s, Category: %s\n", book.Title, book.Author, book.Category)
	}
}

func JsonSerialize(newBook Book) {
	jsonBook := JsonBook{Title: newBook.Title, Author: newBook.Author, Category: newBook.Category}

	jsonData, err := json.Marshal(jsonBook)

	if err != nil {
		fmt.Println("Error when serializing JSON:", err)
		return
	}

	fmt.Println("JsonSerialize: " + string(jsonData))
}

// DeleteBook removes a book from the slice at the specified index.
func DeleteBook(index int) error {
	// Check if the index is within the bounds of the books slice
	if index < 0 || index >= len(books) {
		return fmt.Errorf("Book doesn't exist")
	}

	// Remove the book at the specified index
	books = append(books[:index], books[index+1:]...)

	// Serialize result and print it to console
	ViewBooks()

	return fmt.Errorf("Book removed successfully!")
}

// EditBook updates the book's properties at the specified index.
func EditBook(index int, newTitle, newAuthor, newCategory string) error {
	// Check if the index is within the bounds of the books slice
	if index < 0 || index >= len(books) {
		return fmt.Errorf("index out of bounds")
	}

	// Update the book's properties if the new values pass the validation
	if err := books[index].SetTitle(newTitle); err != nil {
		return err
	}
	if err := books[index].SetAuthor(newAuthor); err != nil {
		return err
	}
	if err := books[index].SetCategory(newCategory); err != nil {
		return err
	}

	// Serialize result and print it to console
	ViewBooks()

	return fmt.Errorf("Book edited successfully!")
}
