package main

import "fmt"

// Function returning a closure
func counter() func() int {
	count := 0 // Enclosed variable
	x1 := 100
	// Just use it temporarily.
	x1++

	return func() int {
		count++ // Modifies outer variable
		return count
	}
}

func main() {
	next := counter()   // next is a closure
	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3

	newCounter := counter()   // New closure with a separate count variable
	fmt.Println(newCounter()) // 1
}
