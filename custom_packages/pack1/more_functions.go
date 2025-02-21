package pack1

import "fmt"

var My_var int = 20

func init() {
	fmt.Println("This is init() called from pack1.more_functions.go")
}
func Print_var() {
	fmt.Printf("The shared variable is %d", My_var)
}
func More_hello() {
	fmt.Println("Hello")
}
