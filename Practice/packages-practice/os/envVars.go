package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

func main() {
	fmt.Println(debug.SetGCPercent(70))
	fmt.Println(debug.SetGCPercent(70))
	data := os.Setenv("GOGC", "200")
	fmt.Println(data)
	fmt.Println(debug.SetGCPercent(170))

}
