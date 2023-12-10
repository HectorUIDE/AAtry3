package main

import (
	"AAtries/Book"
	"AAtries/Category"
	"AAtries/File"
	"AAtries/Search"
	"fmt"
	"os"
)

func main() {

	for {
		fmt.Println("Electronic Book Management System...")
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
			newBook1 := Book.Book{}

			var title string
			fmt.Print("Enter Title: ")
			fmt.Scanln(&title)
			fmt.Println(newBook1.SetTitle(title))

			var author string
			fmt.Print("Enter Author: ")
			fmt.Scanln(&author)
			fmt.Println(newBook1.SetAuthor(author))

			var category string
			fmt.Print("Enter Category: ")
			fmt.Scanln(&category)
			fmt.Println(newBook1.SetCategory(category))

			fmt.Println(newBook1.AddBook())
		case 2:
			Book.ViewBooks()
		case 3:
			newCategory1 := Category.Category{}

			var category string
			fmt.Print("Enter Category Name: ")
			fmt.Scanln(&category)
			fmt.Println(newCategory1.SetCategory(category))

			fmt.Println(newCategory1.AddCategory())
		case 4:
			Category.ViewCategories()
		case 5:
			searchMenu()
		case 6:
			File.SaveToFile()
		case 7:
			File.LoadFromFile(Book.GetBooksPointer())
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
		Search.SearchByTitle(title)
	case 2:
		var author string
		fmt.Print("Enter Author to search: ")
		fmt.Scanln(&author)
		Search.SearchByAuthor(author)
	case 3:
		var category string
		fmt.Print("Enter Category to search: ")
		fmt.Scanln(&category)
		Search.SearchByCategory(category)
	default:
		fmt.Println("Invalid search option. Returning to main menu.")
	}
}

/*
//MODULE_1_BOOK---------------------------------------------------------------------------------------------

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

//MODULE_2_CATEGORY---------------------------------------------------------------------------------------

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

//MODULE_3_SEARCH------------------------------------------------------------------------------------------------------

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

//MODULE_4_FILE----------------------------------------------------------------------------------------------------

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
*/
