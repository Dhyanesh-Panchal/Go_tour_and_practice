package main

import "fmt"

func main() {
	// Integer Conversions to Unsigned int
	var x int32 = 1000
	var y uint8 = uint8(x) // returns the walkthrough value 256 -> 255 + 1 -> 0 (for uint8)
	fmt.Println(y)
	// Int to uint

	var x1 int8 = -1
	var x2 = uint64(x1)
	fmt.Println("int8 to uint8", x2)

	// Float to Int
	var f float64 = 129.9
	var i int8 = int8(f) // Drops decimal (does not round)
	fmt.Println(i)       // Output:  -127 (not 130)
	var f2 float64 = 3.9
	var i2 int = int(f2) // Drops decimal (does not round)
	fmt.Println(i2)      // Output: 3 (not 4)

	// Byte and Rune
	var b byte = 'A' // Works fine
	var r rune = '界'
	fmt.Println(b, r)

	// String to Byte Slice
	s := "Hello"
	bs := []byte(s) // Converts string to []byte
	fmt.Println(bs)

	//var bad_i []int = []int(s) 		THIS will give error
	//fmt.Println(bad_i)

	fmt.Println("------------- type conversion with custom types ---------------")
	type Mytype int
	type Mytype2 int
	type Mytype3 = int

	var v1 Mytype = 10
	//var v2 Mytype2 = Mytype3(5)  THIS WIL CAUSE ERROR
	//var v2 Mytype2 = Mytype3(5)
	//var v3 Mytype3 = v1
	fmt.Println(v1)

	// --------------- rune to string ------------
	arr := []rune{97, 98, 99, -4, -5, -6, 66}

	fmt.Println(arr)
	fmt.Println(string(arr))

}
