package main

import "fmt"

func mera_func(x, y int, z string) int {
	return 0
}

func main() {
	s := []int{1, 2, 3}
	a := [...]int{1, 2, 3}
	s1 := a[:2]
	//s_nil := a[:0]
	//fmt.Println(s_nil == nil)

	fmt.Printf("%T %T %T", a, s, s1)
}
