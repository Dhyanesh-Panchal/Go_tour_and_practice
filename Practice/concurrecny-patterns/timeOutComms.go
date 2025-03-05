package main

import (
	"fmt"
	"time"
)

func tooSlow(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hey!"
}

func main() {
	c := make(chan string)
	go tooSlow(c)
	select {
	case msg := <-c:
		fmt.Println(msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Waited too long. Bye!")
	}
}
