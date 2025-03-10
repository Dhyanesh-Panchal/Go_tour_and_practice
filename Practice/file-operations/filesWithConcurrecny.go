package main

import (
	"fmt"
	"os"
	"time"
)

func reader(file *os.File) {
	fmt.Println("reader start")

	buff := make([]byte, 32)
	n, err := file.Read(buff)
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}

	fmt.Println(string(buff[:n]))

}

func writer(file *os.File) {
	fmt.Println("writer start")
	n, err := file.WriteAt([]byte("hello world"), 0)

	if err != nil {
		fmt.Println("Error writing file: ", err)
	} else {
		fmt.Println(n)
	}

}

func main() {
	fmt.Println("PID:", os.Getpid())
	file, err := os.OpenFile("./dangerFile.txt", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}

	//n, err := file.WriteString("This is some Text")
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println(n)
	//}
	reader(file)

	//writer(file)

	defer file.Close()

	time.Sleep(time.Minute)
}
