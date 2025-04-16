package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.OpenRoot("./")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(f.Mkdir("testDIR", 0777))
}
