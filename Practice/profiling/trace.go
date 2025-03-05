package main

import (
	"os"
	//"runtime"
	"runtime/trace"
	"sync"
)

func SomeWork(b *uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	*b = (*b)*(*b) - 20
}

func useSomeHeap(length int, wg *sync.WaitGroup) {
	defer wg.Done()
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = 'a'
	}
	clear(b)

	//b = nil
	//runtime.GC()
}

func main() {
	var wg sync.WaitGroup
	traceDumpFile, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	trace.Start(traceDumpFile)
	defer trace.Stop()
	var b uint64 = 2
	wg.Add(110)
	for i := 0; i < 10; i++ {
		go SomeWork(&b, &wg)
	}

	for i := 0; i < 100; i++ {
		go useSomeHeap((i+1)*1000, &wg)
	}
	wg.Wait()

}
