package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	for {
		fmt.Println("Electronic Book Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. View Books")
		fmt.Println("3. Add Category")
		fmt.Println("4. View Categories")
		fmt.Println("5. Search Books")
		fmt.Println("6. Save to File")
		fmt.Println("7. Load from File")
		fmt.Println("8. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addBook()
		case 2:
			viewBooks()
		case 3:
			addCategory()
		case 4:
			viewCategories()
		case 5:
			searchMenu()
		case 6:
			saveToFile()
		case 7:
			loadFromFile()
		case 8:
			fmt.Println("Exiting program.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func searchMenu() {
	fmt.Println("Search Options:")
	fmt.Println("1. Search by Title")
	fmt.Println("2. Search by Author")
	fmt.Println("3. Search by Category")

	var searchChoice int
	fmt.Print("Enter your search option: ")
	fmt.Scanln(&searchChoice)

	switch searchChoice {
	case 1:
		var title string
		fmt.Print("Enter Title to search: ")
		fmt.Scanln(&title)
		searchByTitle(title)
	case 2:
		var author string
		fmt.Print("Enter Author to search: ")
		fmt.Scanln(&author)
		searchByAuthor(author)
	case 3:
		var category string
		fmt.Print("Enter Category to search: ")
		fmt.Scanln(&category)
		searchByCategory(category)
	default:
		fmt.Println("Invalid search option. Returning to main menu.")
	}
}

//---------------------------------------------------------------------------------------------

type Book struct {
	Title    string
	Author   string
	Category string
}

var books []Book

func addBook() {
	var newBook Book
	fmt.Print("Enter Title: ")
	fmt.Scanln(&newBook.Title)
	fmt.Print("Enter Author: ")
	fmt.Scanln(&newBook.Author)
	fmt.Print("Enter Category: ")
	fmt.Scanln(&newBook.Category)

	books = append(books, newBook)
	fmt.Println("Book added successfully!")
}

func viewBooks() {
	fmt.Println("Books:")
	for _, book := range books {
		fmt.Printf("Title: %s, Author: %s, Category: %s\n", book.Title, book.Author, book.Category)
	}
}

//---------------------------------------------------------------------------------------

type Category struct {
	Name string
}

var categories []Category

func addCategory() {
	var newCategory Category
	fmt.Print("Enter Category Name: ")
	fmt.Scanln(&newCategory.Name)

	categories = append(categories, newCategory)
	fmt.Println("Category added successfully!")
}

func viewCategories() {
	fmt.Println("Categories:")
	for _, category := range categories {
		fmt.Println(category.Name)
	}
}

//------------------------------------------------------------------------------------------------------

func searchByTitle(title string) {
	fmt.Printf("Search Results for Title '%s':\n", title)
	for _, book := range books {
		if book.Title == title {
			fmt.Printf("Author: %s, Category: %s\n", book.Author, book.Category)
		}
	}
}

func searchByAuthor(author string) {
	fmt.Printf("Search Results for Author '%s':\n", author)
	for _, book := range books {
		if book.Author == author {
			fmt.Printf("Title: %s, Category: %s\n", book.Title, book.Category)
		}
	}
}

func searchByCategory(category string) {
	fmt.Printf("Search Results for Category '%s':\n", category)
	for _, book := range books {
		if book.Category == category {
			fmt.Printf("Title: %s, Author: %s\n", book.Title, book.Author)
		}
	}
}

//----------------------------------------------------------------------------------------------------

const fileName = "books.json"

func saveToFile() error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(books)
	if err != nil {
		return err
	}

	return nil
}

func loadFromFile() error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&books)
	if err != nil {
		return err
	}

	return nil
}
