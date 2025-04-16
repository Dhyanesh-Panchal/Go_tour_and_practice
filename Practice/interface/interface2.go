package main

import (
	"fmt"
	"reflect"
)

type myStruct struct {
	val int
}

type Inter1 interface {
	AddIt(int)
}

type Inter2 interface {
	AddIt(int)
}

func (x myStruct) AddIt(y int) {
	x.val += y
	fmt.Println(reflect.TypeOf(x))
}

func main() {
	s1 := &myStruct{
		val: 3,
	}
	s2 := &myStruct{
		val: 3,
	}
	var i1 Inter1 = s1

	var i2 Inter2 = s2

	fmt.Println(i1, reflect.TypeOf(i1), i2, reflect.TypeOf(i2))
	fmt.Println(i1 == i2)

}
