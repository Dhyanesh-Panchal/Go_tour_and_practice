package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func newFunc(id int, wg *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()
	defer fmt.Println("Defer of newFunc")
	defer wg.Done()
	defer fmt.Println("Defer of newFunc #2")

	for {
		time.Sleep(3 * time.Second)
		fmt.Println("Exitting #", id)
		runtime.Goexit()
		fmt.Println("After the exit")
	}
}

func main() {
	var wg sync.WaitGroup
	fmt.Println(runtime.GOMAXPROCS(0))
	wg.Add(3)
	go newFunc(1, &wg)
	go newFunc(2, &wg)
	go newFunc(3, &wg)
	//runtime.Goexit() This will terminate the main call, wait for all routines to complete, and then will panic as soon as
	fmt.Println("Waiting......")
	fmt.Println("Number of Go routines: ", runtime.NumGoroutine())
	wg.Wait()
	time.Sleep(1 * time.Second)
}
