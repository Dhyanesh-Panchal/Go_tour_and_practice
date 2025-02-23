package main

import "fmt"

func main() {
	var printer_fn []func()

	var x = [4]int{1, 2, 3, 4}
	for _, v := range x {
		printer_fn = append(printer_fn, func() {
			fmt.Printf("\n%v\n", v)
		})
	}

	for _, fn := range printer_fn {
		fn()
	}
}
