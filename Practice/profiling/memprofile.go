package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"strings"
)

func memoryHeavy(s string, repeatFactor int) string {
	newStr := strings.Repeat(s, repeatFactor)
	return newStr
}

func main() {

	f, err := os.Create("mem.prof")
	if err != nil {
		panic(err)
	}

	// Track process heap
	pprof.WriteHeapProfile(f)
	defer f.Close()

	memoryHeavy("Some Text", 1000)
	s := memoryHeavy("Some more Text", 10000)

	fmt.Println(s)

}
