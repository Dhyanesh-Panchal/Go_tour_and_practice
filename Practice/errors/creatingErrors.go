package main

import (
	"errors"
	"fmt"
)

// Errors

func main() {
	err := errors.New("This is error message")

	// Alternative
	//fmt.Errorf("This is error message")
	fmt.Println(err)

}
