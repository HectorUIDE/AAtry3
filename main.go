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
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Welcome to the Electronic Book Management System!")
		})
		http.HandleFunc("/addbook", addBookHandler)
		http.HandleFunc("/viewbooks", viewBooksHandler)
		http.HandleFunc("/addcategory", addCategoryHandler)
		http.HandleFunc("/searchbookstitle", searchBooksByTitleHandler)

		fmt.Println("Server is running on port 8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	*/
	for {
		fmt.Println("Electronic Book Management System...")
		fmt.Println("1. Add Book")
		fmt.Println("2. View Books")
		fmt.Println("3. Add Category")
		fmt.Println("4. View Categories")
		fmt.Println("5. Search Books")
		fmt.Println("6. Save to File")
		fmt.Println("7. Load from File")
		fmt.Println("8. Remove book")
		fmt.Println("9. Edit book")
		fmt.Println("10. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1: // Add Book
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
		case 2: // View Books
			Book.ViewBooks()
		case 3: // Add Category
			newCategory1 := Category.Category{}

			var category string
			fmt.Print("Enter Category Name: ")
			fmt.Scanln(&category)
			fmt.Println(newCategory1.SetCategory(category))

			fmt.Println(newCategory1.AddCategory())
		case 4: // View Categories
			Category.ViewCategories()
		case 5: // Search Book
			searchMenu()
		case 6: // Save File
			//			File.SaveToFile()
			// Implement goroutines so we don't have to wait for file SAVING to end
			go func() {
				if err := File.SaveToFile(); err != nil {
					fmt.Println("Error saving to file:", err)
				}
			}()
			fmt.Println("Save initiated...")

		case 7: // Load file
			//			File.LoadFromFile(Book.GetBooksPointer())
			// Implement goroutines so we don't have to wait for file LOADING to end
			go func() {
				if err := File.LoadFromFile(Book.GetBooksPointer()); err != nil {
					fmt.Println("Error loading from file:", err)
				}
			}()
			fmt.Println("Loading initiated...")

		case 8: // Delete
			Book.ViewBooks()
			var index int
			fmt.Print("Enter the number of the book you wish to delete: ")
			fmt.Scanln(&index)
			fmt.Println(Book.DeleteBook(index))
		case 9: // Edit
			Book.ViewBooks()

			var index int
			fmt.Print("Enter the number of the book you wish to update: ")
			fmt.Scanln(&index)

			var title string
			fmt.Print("Enter new Title: ")
			fmt.Scanln(&title)

			var author string
			fmt.Print("Enter new Author: ")
			fmt.Scanln(&author)

			var category string
			fmt.Print("Enter new Category: ")
			fmt.Scanln(&category)

			fmt.Println(Book.EditBook(index, title, author, category))
		case 10: //Exit
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
func addBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var newBook Book.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = newBook.AddBook()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Book added successfully"))
}

func viewBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	books := Book.GetBooks()
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func addCategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var newCategory Category.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = newCategory.AddCategory()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Category added successfully"))
}

func searchBooksByTitleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "Title parameter is required", http.StatusBadRequest)
		return
	}

	results := Search.SearchByTitle(title)
	err := json.NewEncoder(w).Encode(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
*/
