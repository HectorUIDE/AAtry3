package main

/*
import (
	"encoding/json"
	"os"
)

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
