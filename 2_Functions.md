# Functions

```go
func myfn(x int, y,z string) bool{
	// Body
}
```
- types comes after variable names.
- here `y` and `z` are both of type string

```go
// Multiple return.
func swap(x, y string) (string, string) {
	return y, x
}

// Named return
func split(sum int) (x, y int) {
x = sum * 4 / 9
y = sum - x
return
}
```
- A function can return multiple values.
- A Named return, they are treated as variables defined at the top of the function. This is known as a "naked" return.
### Variable lenth arguments in functions
```go
func sum(numbers ...int) int {
	total := 0
	
}
```
- When you pass variable-length arguments (`...int`), Go converts them into a slice (`[]int`).
### Anonynomus functions
```go
result := func(a, b int) int {
        return a + b
    }(3, 4)
```

### function has its own type. And function can be assigned to another Variable as well eg.:
```go
func add(a, b int) int {
    return a + b
}

func main() {
    var operation func(int, int) int // Function type declaration
    fmt.Printf("Type: %T\n", operation) // Output: func(int, int) int
    fmt.Println(operation == nil) // Output: true (zero value)
}
```
> The zero value of a function type is `nil`. **If you call `nil` function, it causes a runtime `panic`.**


# Variables
- `var` can be used to declare variables at package and function level.
- Inside a function, the `:=` short assignment statement can be used in place of a `var` declaration with implicit type.
```go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	const x int = 3
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```
- constants can be declared using `const` keyword.

