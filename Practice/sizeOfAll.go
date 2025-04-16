package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := 10
	fmt.Println("make([]int, 1000): ", unsafe.Sizeof(make([]int, 1000)))
	fmt.Println("make(map[int]int): ", unsafe.Sizeof(make(map[int]int)))
	fmt.Println("[5]int{1,2,3,4,5}: ", unsafe.Sizeof([5]int{1, 2, 3, 4, 5}))
	//fmt.Println("make([...]int,10): ", unsafe.Sizeof(make([...]int, 10)))  // This is an Error!!!
	arr := make([...]int, 10)
	fmt.Println(arr)
	fmt.Println("&x: ", unsafe.Sizeof(&x))
}
