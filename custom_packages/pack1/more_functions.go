package pack1

import "fmt"

func init() {
	fmt.Println("This is init() called from pack1.more_functions.go")
}

func More_hello() {
	fmt.Println("Hello")
}
