package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func subworker() {
	panic("panic from subworker")
}
func someWork() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd in worker:", r)

			// Retrive stacktrace
			stackTrace := make([]byte, 1024)
			n := runtime.Stack(stackTrace, false)

			fmt.Println("stack trace:", string(stackTrace[:n]))
		}
	}()

	fmt.Println("Some worker start")

	subworker() // not a seperate routine call
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered #last", r)
		}
	}()
	defer func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered #2", r)
				fmt.Println(reflect.TypeOf(reflect.TypeOf(r)))
			}
			time.Sleep(4 * time.Second)
		}()

		panic("Panic in defer")

		if r := recover(); r != nil {
			fmt.Println("Recovered :", r)
		}

	}()

	go someWork()
	//go someWork()

	panic("panic from main")

}
