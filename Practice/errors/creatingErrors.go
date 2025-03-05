package main

import (
	"errors"
	"fmt"
)

// Errors

func main() {
	err := errors.New("this is an error")
	fmt.Println(err)
}
