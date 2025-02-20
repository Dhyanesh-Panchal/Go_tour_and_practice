package pack1

import "fmt"

func init() {
	fmt.Println("This is init() called from pack1.package1.go")
}

func Hello() {
	fmt.Println("Hello From Pack1")
}
