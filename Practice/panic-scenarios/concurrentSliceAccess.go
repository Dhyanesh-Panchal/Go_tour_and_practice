package main

import (
	"fmt"
	"time"
)

func concurrentSliceReader(m []int) {
	for i := 0; i < 10; i++ {
		go func() {

			x1 := m[1]
			x2 := m[2]
			fmt.Println(x1, x2)

		}()
	}
}

func concurrentSliceWriter(m []int) {
	time.Sleep(1 * time.Microsecond)
	for i := 0; i < 10; i++ {
		go func() {

			// Write
			m[1] = i % 10
			m[2]++
			_ = m[3]

		}()
	}
}

func main() {
	m := make([]int, 10)

	go concurrentSliceReader(m)
	go concurrentSliceWriter(m)

	time.Sleep(5 * time.Second)
	fmt.Println(m)
}

// Try to run with go run -race concurrentSliceAccess.go
