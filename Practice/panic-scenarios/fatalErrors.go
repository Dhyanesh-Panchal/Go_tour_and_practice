package main

import (
	"fmt"
	"log"
	"time"
)

func fatalFunction() {
	fmt.Println("Inside the fatal function")

	log.Fatal("This is some fatal error")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	go fatalFunction()

	fmt.Println("THis is some code")
	time.Sleep(4 * time.Second)
}
