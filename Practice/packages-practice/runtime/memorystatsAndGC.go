package main

import (
	"fmt"
	"runtime"
)

func getMemoryStats() runtime.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return m
}

func PrintMemoryStats(m runtime.MemStats) {
	fmt.Printf("Alloc = %v \n", m.Alloc)
	fmt.Printf("\tTotalAlloc = %v \n", m.TotalAlloc)
	fmt.Printf("\tSys = %v \n", m.Sys)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func main() {
	
}
