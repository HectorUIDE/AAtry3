package File

import (
	"AAtries/Book"
	"encoding/json"
	"fmt"
	"os"
)

const fileName = "books.json"

func SaveToFile() error {
	books := Book.GetBooks()
	fmt.Printf("Printing Book.GetBooks(): %v\n", Book.GetBooks())
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

func LoadFromFile(books *[]Book.Book) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	// Use the pointer to update the original books slice
	err = decoder.Decode(books)
	if err != nil {
		return err
	}

	return nil
}
