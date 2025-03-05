package main

import "fmt"

func Reverse(s string) string {
	//b := []byte(s)   // Will not work for all Unicode

	b := []rune(s) // Will work for all Unicode

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	
	return string(b)
}

func main() {
	input := "This is the input 世"
	fmt.Println(input)
	fmt.Println(Reverse(input))
	fmt.Println(Reverse(Reverse(input)))
}
