package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println(os.Getpid())

	rfile, err := os.Open("./fileIOReadFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 1024)
	n, err := rfile.Read(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, err)
	fmt.Println(string(data[:n]))

	file, err := os.OpenFile("./MeriFile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte("This is text written by me."))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Write success!")
	}

	// Reading with a buffered reader.
	fmt.Println("-------- Reading with buffered reader --------")
	reader := bufio.NewReader(rfile)

	//buff := make([]byte, 100) Use this if we need custom buffer
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				fmt.Println("Reached EOF")
				break
			}
			panic(err)
		}

		fmt.Print(line)
	}

	f, err := os.CreateTemp("", "example.txt")
	if err != nil {

		panic(err)
	}

	n, err = f.Write([]byte("This is text written by me."))
	if err != nil {
		panic(err)
	}
	fmt.Println(f.Name())
	time.Sleep(3 * time.Second)
	reader = bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println(line)
				break
			}
		}
		fmt.Println(line)
	}
	os.Remove(f.Name())

	f.Close()

	// Error case : read /path/to/file: file already closed.
	//n, err = f.Read(data)
	//if err != nil {
	//	fmt.Println("Error reading Closed file: ", err)
	//}
	//fmt.Println(string(data[:n]))
}
