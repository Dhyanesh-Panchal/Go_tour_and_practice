package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

var mu sync.Mutex

func waiting() {
	mu.Lock()
	time.Sleep(2 * time.Second)
	mu.Unlock()
}

func main() {
	blockProfF, err := os.Create("block.prof")
	if err != nil {
		fmt.Println(err)
	}
	defer blockProfF.Close()

	mutexProfF, err := os.Create("mutex.prof")
	if err != nil {
		fmt.Println(err)
	}
	defer mutexProfF.Close()

	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)

	for i := 0; i < 10; i++ {
		go waiting()
	}

	time.Sleep(3 * time.Second)
	pprof.Lookup("block").WriteTo(blockProfF, 0)
	pprof.Lookup("mutex").WriteTo(mutexProfF, 0)

}
