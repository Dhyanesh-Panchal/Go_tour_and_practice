package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	b.WriteString("This is some content")
	fmt.Println(b.Cap())

	fmt.Println(b.String())

	var myrune rune = rune('好')
	fmt.Println(b.WriteRune(myrune)) // returns 3,<nil>
	fmt.Println(b.String())

}
