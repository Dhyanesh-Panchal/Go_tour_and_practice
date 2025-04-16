package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func memUsage() uint64 {
	var m runtime.MemStats
	runtime.GC() // Force garbage collection for accurate measurement
	runtime.ReadMemStats(&m)
	return m.Alloc
}

func main() {
	before := memUsage()
	ch1 := make(chan int) // Unbuffered
	after := memUsage()
	//fmt.Println("Unbuffered channel memory:", after-before, "bytes")
	fmt.Println("Unbuffered channel memory:", unsafe.Sizeof(ch1), "bytes")

	before = memUsage()
	ch2 := make(chan int, 1000) // Buffered with capacity 1000
	after = memUsage()
	fmt.Println("Buffered channel (1000) memory:", after-before, "bytes")

	_ = ch1
	_ = ch2
}
