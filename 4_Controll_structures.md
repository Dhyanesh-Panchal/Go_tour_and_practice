# Control Structures

## For
```go
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```
Go has only one looping construct, the for loop.

The basic for loop has three components separated by semicolons:

- `init` statement: executed before the first iteration
- `condition` expression: evaluated before every iteration
- `post` statement: executed at the end of every iteration

- The `init` and `post` statements are optional. And the for becomes `while` loop!
```go
func main() {
	sum := 1
	for sum < 1000 { // Go's while equivalent.
		sum += sum
	}
	fmt.Println(sum)
}
```

The loop will stop iterating once the boolean condition evaluates to false.

Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.

## If
```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
	)
}
```

- The `if` statement can start with a short statement to execute before the condition.
- Variables declared by the statement are only in scope until the end of the `if` or any following `else`.

## Switch
```go
import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```
- Note here there is not need of `break` 
- Another important difference is that Go's `switch` cases need not be constants, and the values involved need not be integers.
- Switch without a condition is the same as `switch true`.

# `defer`
```go
func main() {
x:=10

defer fmt.Printf("world %d",x)

x = 100

fmt.Println("hello")
}

// The Output is:
hello world 10
```
- A `defer` statement defers the execution of a function until the surrounding function returns. 
- The **deferred call's arguments are evaluated immediately**, but the function call is not executed until the surrounding function returns.
- Multiple Defer callss are stacked in LIFO fashion. Eg:
```go
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d ",i)
	}

	fmt.Println("done")
}
/* Output: 
counting
done 
9 8 7 6 5 4 3 2 1
*/
```


# Range
The range form of the for loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a **copy** of the element at that index.

