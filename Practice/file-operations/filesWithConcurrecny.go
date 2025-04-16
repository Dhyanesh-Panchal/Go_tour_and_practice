package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var readyFlag = false

func main() {
	fmt.Println("PID:", os.Getpid())
	file, err := os.OpenFile("./dangerFile.txt", os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	//n, err := file.WriteString("This is some Text")
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println(n)
	//}

	ready := sync.NewCond(&sync.Mutex{})

	go writer(1, file, ready)
	go writer(2, file, ready)

	go reader(file, ready)
	go reader(file, ready)

	time.Sleep(10 * time.Second)
}

func reader(file *os.File, ready *sync.Cond) {
	if !readyFlag {
		ready.L.Lock()
		ready.Wait()
	}

	fmt.Println("reader start")

	for range time.Tick(time.Second) {
		buff := make([]byte, 32)
		n, err := file.Read(buff)
		if err != nil {

			fmt.Println("Error reading file: ", err)
		}

		fmt.Println("Reader Read: ", string(buff[:n]))

		// Alternate Option
		//data, err := io.ReadAll(file)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Println("Reader Read:", string(data))
	}

}

func writer(id int, file *os.File, ready *sync.Cond) {

	if !readyFlag {
		ready.L.Lock()
		readyFlag = true
		ready.L.Unlock()
		ready.Broadcast()
	}
	fmt.Println("writer start")

	msg := []byte(fmt.Sprintf("Hello From writer #%d\n", id))

	for range time.Tick(time.Second) {
		n, err := file.Write(msg)
		if err != nil {
			fmt.Println("Error writing file: ", err)
		} else {
			fmt.Println("wrote: ", n)
		}
	}

}
