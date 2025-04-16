package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
)

func main() {
	traceStop := startTracing()
	defer traceStop()

	memProfileFile, err := os.Create("./packages-practice/runtime/mem_2.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer memProfileFile.Close()

	m := getMemoryStats()

	printMemoryStats(m)

	// Use some memory
	var s []int
	for i := range 100 {
		s = nil
		s = make([]int, 1024*10)
		for i := 0; i < 1024*10; i++ {
			s[i] = i
		}
		if i%10 == 0 {
			runtime.ReadMemStats(&m)
			printMemoryStats(m)
		}

	}
	runtime.GC()
	pprof.WriteHeapProfile(memProfileFile)

}

func getMemoryStats() runtime.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return m
}

func printMemoryStats(m runtime.MemStats) {
	fmt.Printf("Alloc = %v KB\n", m.Alloc/1024)
	fmt.Printf("\tTotalAlloc = %v KB\n", m.TotalAlloc/1024)
	fmt.Printf("\tSys = %v KB\n", m.Sys/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func startTracing() func() {
	traceDumpFile, err := os.Create("./packages-practice/runtime/trace.out")
	if err != nil {
		log.Fatal(err)
	}

	err = trace.Start(traceDumpFile)

	if err != nil {
		log.Fatalf("Error starting tracing: %v\n", err)
	}

	return func() {
		fmt.Println("Tracing exited")
		trace.Stop()
		err := traceDumpFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}
