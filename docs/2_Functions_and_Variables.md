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
#### Expanding a Slice (... Operator)
We can pass a slice to a variadic function:
```go
numbers := []int{1, 2, 3}
sum(numbers...) // Expands slice into individual arguments

```
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

### functions can be used and passed as a values.

```go
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```

## Closures
- Closures are excellent way of retaining the Variables of the parent function.
- Its works by returning the function which accesses that variable.
- A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is **bound** to the variables.
```go
// Function returning a closure
func counter() func() int {
    count := 0  // Enclosed variable
    var x1 = 100
	// Utilize this x1
    return func() int {
        count++  // Modifies outer variable
        return count
    }
}

func main() {
next := counter()  // next is a closure
fmt.Println(next()) // 1
fmt.Println(next()) // 2
fmt.Println(next()) // 3

newCounter := counter() // New closure with a separate count variable
fmt.Println(newCounter()) // 1
}
```

- The variabe `count` here is stored initially on call stack, but when the counter() returns, this variable **escapes to heap**.
- Note: the `x1` is not used in the return function hence, its get wipedout with stack when call returns.

# Methods
- Methods are functions defined on types.
```go
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
```
- methods can be defined **only on localy defined types**.
---
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
- Overflow is checked while initializing a variable/constant:
- ``const x uint8 = 300 // Compilation Error``
- Also Assigning `nil` to non pointer types ca also cause error:
- ``var x int = nil`` (its not a zero value of `int`)
### Numeric constants
In Go, numeric constants represent fixed values of numbers. They can be integers, floating-point numbers, or complex numbers.

Numeric constants are high-precision values.

Go numeric constants can be either untyped or typed:
- Untyped Constants: Do not have an explicit type declared. Their type is inferred from the context in which they are used. This allows them to be used with different numeric types without explicit conversion.


- Typed Constants: Have an explicit type declared. They can only be used with values of the same type or values that can be implicitly converted to that type.
```go
const hello = "Hello!"
const typedHello string = "Hello!!"
```
- hello is also an untyped string constant. An untyped constant is just a value, one not yet given a defined type that would force it to obey the strict rules that prevent combining differently typed values.

For More details read: [This](https://go.dev/blog/constants#:~:text=This%20is%20an%20untyped%20string%20constant%2C)


