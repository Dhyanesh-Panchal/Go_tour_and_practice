package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan int, 10)

	// Two Listeners
	go func() {
		// Reads from the channel
		for d := range channel {
			fmt.Println("#one: ", d)
		}
	}()

	go func() {
		// Reads from the channel
		for d := range channel {
			fmt.Println("#two: ", d)
		}
	}()

	for i := 0; i < 1000; i++ {
		channel <- i
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	}
	close(channel)
	fmt.Println("Channel closed")

}
