package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {

	rfile, err := os.Open("./fileIOReadFile.txt")
	if err != nil {
		panic(err)
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
	reader := bufio.NewReader(rfile)

	buff := make([]byte, 100)
	for {
		n, err := reader.Read(buff)

		if err != nil {
			if err == io.EOF {
				fmt.Println("Reached EOF")
				break
			}
			panic(err)
		}

		fmt.Print(string(buff[:n]))
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
		n, err := reader.Read(buff)
		if err != nil {
			if err == io.EOF {
				fmt.Println(string(buff[:n]))
				break
			}
		}
		fmt.Println(string(buff[:n]))
	}

	f.Close()

}
