package main

import "fmt"

func main() {
	defer fmt.Println("After the recover")

	defer func() {
		// Catch the panic from the current routine stack
		if r := recover(); r != nil {
			fmt.Println("panic catched", r)
		}
	}()

	defer func() {
		panic("New") // this will override the previous panic
	}()

	s := make([]int, 100)

	fmt.Println(s[4000]) // this will cause panic of type runtime error
}
