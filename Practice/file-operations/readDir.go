package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	files, err := os.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.Type(), file.IsDir())
		fmt.Println(file.Info())

		println(file.IsDir())
	}
}
