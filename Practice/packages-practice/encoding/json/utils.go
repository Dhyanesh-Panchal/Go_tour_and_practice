package main

import (
	"log"
	"os"
)

func FetchData(filePath string) ([]byte, error) {
	data, error := os.ReadFile(filePath)
	if error != nil {
		log.Fatal(error)
		return nil, error
	}

	return data, nil
}
