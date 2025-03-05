package main

import (
	"fmt"
	"time"
)

func concurrentReader(m map[string]string) {

	for i := 0; i < 1; i++ {
		go func() {
			// read,  will work Only if there ar NO concurrent write on going
			// NOTE: This Works on key level, if we are reading a key and other routine is writing the same key, then only there will be panic.
			_ = m["a"]
			_ = m["newKey"]

		}()
	}
}

func concurrentWriter(m map[string]string) {

	for i := 0; i < 100; i++ {
		go func() {
			// This will not catch as: THE ERROR IS FATAL, will crash the program
			defer func() {
				if r := recover(); r != nil {
					fmt.Println(r)
				}
			}()

			// Write
			m["X"] = "x"
			_ = m["X"]

		}()
	}
}

func main() {
	m := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
	}

	go concurrentWriter(m)
	go concurrentReader(m)

	time.Sleep(5 * time.Second)
}
