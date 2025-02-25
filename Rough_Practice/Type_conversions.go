package main

import "fmt"

func main() {
	// Integer Conversions to Unsigned int
	var x int32 = 256
	var y uint8 = uint8(x) // returns the walkthrough value 256 -> 255 + 1 -> 0 (for uint8)
	fmt.Println(y)

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

}
