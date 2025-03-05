package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	// some work
	time.Sleep(1 * time.Second)

	panic("panic from worker" + string(id))
}

func main() {
	// This will not catch the panic from different routines.
	// Defer
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Catched in main: ", r)
		}
		return
	}()

	go worker(1)
	go worker(2)

	time.Sleep(4 * time.Second)
	fmt.Println("done")

}
