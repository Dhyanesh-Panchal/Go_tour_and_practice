package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println(i)
		}()
	}
	time.Sleep(2 * time.Second)
}
