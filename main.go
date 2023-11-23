package main

import (
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
			AddBook()
		case 2:
			ViewBooks()
		case 3:
			//			addCategory()
		case 4:
			//			viewCategories()
		case 5:
			//			searchMenu()
		case 6:
			//			saveToFile()
		case 7:
			//			loadFromFile()
		case 8:
			fmt.Println("Exiting program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
