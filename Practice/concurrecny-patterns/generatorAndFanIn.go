package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Generator(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			c <- msg
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}

func FanIn(inputs ...<-chan string) <-chan string {
	outputChan := make(chan string)

	for _, inputChan := range inputs {
		go func() {
			for i := 0; i < 10; i++ {
				outputChan <- <-inputChan
			}
		}()
	}
	return outputChan
}

func main() {
	generatorCh1 := Generator("Hey 1!")
	generatorCh2 := Generator("Hey 2!")

	fanInCh1 := FanIn(generatorCh1, generatorCh2)
	for i := 0; i < 20; i++ {
		fmt.Println(<-fanInCh1)
	}
}
