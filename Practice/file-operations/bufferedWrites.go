package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const chunkSize int = 1024

func main() {
	data := generateLargeData()
	writeData(data, "./file-operations/largetext.txt")
}

func writeData(data []byte, fp string) {
	file, err := os.Create(fp)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	buffWriter := bufio.NewWriter(file)

	for i := 0; i < len(data)/chunkSize; i++ {
		end := (i + 1) * chunkSize
		if end > len(data) {
			end = len(data)
		}

		_, err := buffWriter.Write(data[i*chunkSize : end])
		if err != nil {
			log.Fatal(err)
		}
		buffWriter.Flush()
	}
}

func generateLargeData() []byte {
	baseStr := "This is some Text\n"

	bigData := strings.Repeat(baseStr, 1024*1024*40)

	return []byte(bigData)
}
