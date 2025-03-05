package main

import (
	"fmt"
	"sync"
	"time"
)

type Resource struct {
	data int
	mu   sync.Mutex
}

// Circular wait
func f(id int, r1, r2 *Resource, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(id, " stated")

	// Lock first resource
	r1.mu.Lock()

	fmt.Println(id, " acquired 1st lock")

	// do some work on it
	time.Sleep(1 * time.Second)

	fmt.Println(id, " Waiting for 2nd lock")

	// simultaniously demand lock on 2nd resource
	r2.mu.Lock()

	_ = r1.data + r2.data

	r2.mu.Unlock()

	r1.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	re1 := Resource{data: 1}
	re2 := Resource{data: 2}

	wg.Add(2)
	go f(1, &re1, &re2, &wg)
	go f(2, &re2, &re1, &wg)

	wg.Wait()

}
