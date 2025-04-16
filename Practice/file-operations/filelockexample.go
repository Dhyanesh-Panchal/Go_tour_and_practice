package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"strconv"
	"time"
)

func lockFile(f *os.File) error {
	return unix.Flock(int(f.Fd()), unix.LOCK_EX)
}

func unlockFile(f *os.File) error {
	return unix.Flock(int(f.Fd()), unix.LOCK_UN)
}

func main() {
	go writeFile(1)
	go writeFile(2)
	time.Sleep(1 * time.Minute)
}

func writeFile(id int) {
	file, err := os.OpenFile("./synchronisedFile.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	fmt.Println("File read successful, for worker:", id)
	if err != nil {
		fmt.Println("This is worker:", id)
		fmt.Println("Error opening file: ", err)
	}
	if err := lockFile(file); err != nil {
		fmt.Println("Error locking file: ", err)
	}

	// Cleanup
	defer func() {
		fmt.Println("Unlocking file: ", file.Name())
		if err := unlockFile(file); err == nil {
			if err := file.Close(); err != nil {
				fmt.Println("Error closing file: ", err)
			}
		} else {
			fmt.Println("Error unlocking file: ", err)
		}
	}()

	n, err := file.Write([]byte("This is Secret data from Process: " + strconv.Itoa(os.Getpid()) + "\n"))
	if err != nil {
		fmt.Println("Error writing file: ", err)
	} else {
		fmt.Println("Wrote ", n, "bytes")
	}
}
