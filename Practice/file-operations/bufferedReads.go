package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./file-operations/largetext.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffReader := bufio.NewReader(file)

	buffer := make([]byte, 1024*1024+4)

	for i := 0; ; i++ {
		n, err := buffReader.Read(buffer)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println("Data #", i)
		fmt.Println(len(buffer[:n]))
	}
}
