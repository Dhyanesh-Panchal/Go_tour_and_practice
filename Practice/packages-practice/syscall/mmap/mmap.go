package main

import (
	"fmt"
	"os"
	"sync"
	"syscall"
	"time"
)

func GetMemoryMappingOfFile(filePath string, size int64) ([]byte, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = file.Truncate(size)
	if err != nil {
		return nil, err
	}

	mapping, err := syscall.Mmap(int(file.Fd()), 0, int(size), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		return nil, err
	}
	return mapping, nil
}

func Writer(mapping []byte, offset int, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 3 {
		time.Sleep(1 * time.Second)
		copy(mapping[offset:], []byte("Concurrently written\n"))

		offset += 30
	}

}

func main() {
	filePath := "mmap_testfile.txt"
	size := int64(5 * 1024 * 1024) // 100MB

	mappedData, err := GetMemoryMappingOfFile(filePath, size)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go Writer(mappedData, 300*i, &wg)
	}

	wg.Wait()

}
