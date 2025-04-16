package main

import "fmt"

func main() {
	defer func() {
		defer func() {
			fmt.Println("2nd defer is called")
			if r := recover(); r != nil {
				fmt.Println("Recovered Error:", r)
			} else {
				fmt.Println("Recover is nil in 2nd defer")
			}
		}()
		fmt.Println("Inside the last defer")
	}()

	defer fmt.Println("Something something")

	fmt.Println("Start")

	panic("This is panic")
}
