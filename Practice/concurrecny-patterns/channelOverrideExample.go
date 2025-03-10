package main

import (
	"fmt"
	"time"
)

func listner() chan int {
	ch := make(chan int, 1)

	go func(ch chan int) int {

		time.Sleep(3 * time.Second)
		fmt.Println("Starting the listen")

		for i := range ch {
			fmt.Println(i)
		}
		return 10
	}(ch)
	//fmt.Println(unsafe.Pointer(&ch))

	return ch
}

// ------------------- This code will deadlock ----------
func main() {
	ch := listner()

	//fmt.Println(unsafe.Pointer(&ch))
	ch = make(chan int)

	for i := 0; i < 10; i++ {
		ch <- i + 1
	}
}
