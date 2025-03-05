package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str any = "Hello Mitro"

	//readonlyValStr := reflect.ValueOf(str)

	writableValStrPtr := reflect.ValueOf(&str)
	writableValStrPtr.Elem().Set(reflect.ValueOf(100))

	fmt.Println(writableValStrPtr.Elem())
}
