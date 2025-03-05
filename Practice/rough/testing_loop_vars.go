package main

import "fmt"

func main() {
	var printerFn []func()

	var x = [4]int{1, 2, 3, 4}
	for _, v := range x {
		printerFn = append(printerFn, func() {
			fmt.Printf("\n%v\n", v)
		})
	}

	for _, fn := range printerFn {
		fn()
	}
}
