package main

import (
	"fmt"
	"unsafe"
)

type MyStruct struct {
	b   int8
	a   int8
	c   int32
	str string
	d   float64
	e   bool
}

func main() {
	s := MyStruct{
		a:   1,
		b:   2,
		str: "hello",
		c:   3,
		d:   3.141592,
		e:   true,
	}
	//x := (*int)((unsafe.Pointer(&s)).)
	fmt.Printf("\n%T, Offsets: a:%d, b:%d, str:%d, c: %d, d: %d, e: %d", s, unsafe.Offsetof(s.a), unsafe.Offsetof(s.b), unsafe.Offsetof(s.str), unsafe.Offsetof(s.c), unsafe.Offsetof(s.d), unsafe.Offsetof(s.e))
	fmt.Println("Total Size:", unsafe.Sizeof("Hellos"))

	ptr := unsafe.Pointer(&s)

	// Accessing the data of field c in s through pointers.
	ptr = unsafe.Add(ptr, unsafe.Offsetof(s.c)) // Pointer Arithmetic
	data := (*int64)(ptr)
	fmt.Println(*data)

	// lets modify
	*(*int64)(ptr) = 100
	fmt.Println(*data)

	// Playing with unsafe type conversions on string
	str := "Hello this is some content"

	// unsafe.StringData() Return the pointer to the underlying (IMMUTABLE) byte array.
	// 		& unsafe.Slice() prepares a slice starting from location given by ptr till the given length.
	bytearr := unsafe.Slice(unsafe.StringData(str), len(str))
	fmt.Printf("%v \n", bytearr)

}
