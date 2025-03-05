package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	f1()
	time.Sleep(3 * time.Second)
	fmt.Println("Exiting the main.")

}

func f1() {
	go f3()
	go f2()
}
func f2() {
	f4()
}

func f3() {
	f5()
}

func f4() {
	time.Sleep(1 * time.Second)
	stackTrace := make([]byte, 1024)

	n := runtime.Stack(stackTrace, true)

	fmt.Println("StackTrace from f4:")
	fmt.Println(string(stackTrace[:n]))

	fmt.Println("stack trace printed!")
}

func f5() {
	time.Sleep(1 * time.Second)
	fmt.Println("f5")
	//debug.PrintStack()
}
