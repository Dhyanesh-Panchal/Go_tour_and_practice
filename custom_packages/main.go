package main

import (
	pack1 "custom_package/pack1"
	pack2 "custom_package/package2"
)

func main() {
	pack1.Hello()
	pack2.Hello()

	//fmt.Printf("\n\nType of a Package is %T", pack1)
	// Above line gives, compile time error,since pack1 is a package name and not a variable, function, or value,
}
