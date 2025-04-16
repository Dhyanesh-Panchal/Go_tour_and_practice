package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"unsafe"
)

func main() {
	//str := "This is some string"

	//buf := make([]byte, 10)
	file, err := os.OpenFile("./test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(unsafe.Sizeof(int(4)))

	err = binary.Write(file, binary.LittleEndian, []byte("1000"))
	if err != nil {
		log.Fatal(err)
	}

}

/*
Go’s binary package provides interface binary.ByteOrder

type ByteOrder interface {
    Uint16([]byte) uint16
    Uint32([]byte) uint32
    Uint64([]byte) uint64
    PutUint16([]byte, uint16)
    PutUint32([]byte, uint32)
    PutUint64([]byte, uint64)
    String() string
}

This interface is implemented by types binary.LittleEndian and binary.BigEndian

*/
