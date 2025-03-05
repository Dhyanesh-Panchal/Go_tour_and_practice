package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}

	m.Store("Key1", 100)
	m.Store("Key2", 200)
	m.Store("Key3", 300)

	val, ok := m.Load("Key1")
	fmt.Println(val, ok)
	val, ok = m.LoadAndDelete("Key2")
	val, ok = m.Load("Key2")
	fmt.Println(val, ok)

	val2, _ := m.Swap("Key1", 1000)
	fmt.Println("Swapped Value from Key1", val2)
	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

}
