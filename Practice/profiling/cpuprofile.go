package main

import (
	"os"
	"runtime/pprof"
)

func SomeWork(b *uint64) {
	*b = (*b)*(*b) - 20
}

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	//var meraBull bool = false
	var meraBull uint64 = 10

	for i := 0; i < 1e9; i++ {
		SomeWork(&meraBull)
	}

	// At last run `go tool pprof cpu.prof`
}
