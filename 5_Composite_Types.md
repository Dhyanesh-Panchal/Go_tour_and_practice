# Composite Types
- Go supports 8 different composite types as specified by the language specification namely `array`, `struct`, `pointer`, `function`, `interface`, `slice`, `map`, and `channel` types.

# Pointers
A pointer holds the memory address of the value.

Addressing(using &) and Dereferencing(using *) is similar to that of C.
```go
func main() {
    i, j := 42, 2701
    
    p := &i         // point to i
    fmt.Println(*p) // read i through the pointer
    *p = 21         // set i through the pointer
    fmt.Println(i)  // see the new value of i
    
    p = &j         // point to j
    *p = *p / 37   // divide j through the pointer
    fmt.Println(j) // see the new value of j
}
```

# Structs
A `struct` is a Collection of fields. '.' can be used to access the fields (Also for pointers to a struct.)
```go
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
```
- To access the field X of a struct when we have the struct pointer `p`, we could write `(*p).X`. However, that notation is cumbersome, so the language permits us to write just `p.X`, without the explicit dereference.
- A struct literal denotes a newly allocated struct value by listing the values of its fields. 
- You can list just a subset of fields by using the `Name:` syntax. (And the order of named fields is irrelevant.)
- The special prefix & returns a pointer to the struct value.
## Anonymous struct
```go
// Declare and initialize an anonymous struct
person := struct {
    Name string
    Age  int
}{
    Name: "John Doe",
    Age:  30,
}
```
### Type Aliasing and Type Indentity: [ref](https://go.dev/ref/spec#Type_identity)
---

# Array
```go
func main() {
	var a [2]string // initialized ["",""]
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```

- Length of array can't be changed, fixed sized.

## Slice
- An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
```go
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
```
- **Slices are like References to underlying array**
- Slices can be directly declared as: ``[]T{...elements}``
- A slice has both a length and a capacity.
    - The `length` of a slice is the number of elements it contains. 
    - The `capacity` of a slice is the number of elements in the underlying array, counting from the first element in the slice. 
    - The length and capacity of a slice s can be obtained using the expressions `len(s)` and `cap(s)`. 
    - You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
```go
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4] 
	printSlice(s) 

	// Drop its first two values.
	s = s[2:] //
	printSlice(s)
	// IMPORTANT NOTE: Reslicing from Front Actually reduces the capacity.
}
/*
Output:
   len=6 cap=6 [2 3 5 7 11 13]
   len=0 cap=6 []
   len=4 cap=6 [2 3 5 7]
   len=2 cap=4 [5 7]
*/

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

- **slicing with range greater than capacity causes panic**
```go
// Extend its length.
s = s[:10] // capacity of this s is 6
printSlice(s) 
```
Causes:
```bash
panic: runtime error: slice bounds out of range [:10] with capacity 6

goroutine 1 [running]:
main.main()
	/tmp/sandbox1254776560/prog.go:14 +0x7b
```

- **Zero Value** of slice is nil.
  - ``var s []int`` has a cap:0 len:0 and `s == nil` is true.
- Slices can be created **dynamically** using `make()`.
```go
a := make([]int, 5)  // len(a)=5
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

### Appending to a slice
- `append()` can be used to append to a slice.
```go
func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	// If len = capacity, appends doubles the capacity and inserts the element.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	// If even after doubling the capacity, if thats not sufficient; append Appends the data and capacity grows one more than length.
	s = append(s, 2, 3, 4, 5, 6, 7, 8)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/*
Output:
    len=0 cap=0 []
    len=1 cap=1 [0]
    len=2 cap=2 [0 1]
    len=9 cap=10 [0 1 2 3 4 5 6 7 8]
*/
```
---
# Maps
A map maps keys to values.
```go
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
```

- Map hashes the keys to store and access the data.
- The zero value of a map is `nil`. A `nil` map has no keys, **nor can keys be added**. Adding a key to `nil` map will cause **panic**.
- The `make` function returns a `map` of the given type, initialized and ready for use.
  Delete an element:

  ``delete(m, key)``

  - Test that a key is present with a two-value assignment:

    ``elem, ok = m[key]``
  - If `key` is in `m`, `ok` is `true`. If not, `ok` is `false`.
  - If `key` is not in the map, then `elem` is the zero value for the map's element type.

---
# Interface
An interface type is defined as a set of method signatures.
- A type implements an interface by implementing its methods.
- A value of interface type can hold any value that implements those methods.
Good Example:
```go
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v // Hence this line will give error

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```
- An interface value holds a value of a specific underlying concrete type.
- Calling a method on an interface value executes the method of the same name on its underlying type.
- **If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.**
- In Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)
```go
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
  var i I
  var t *T
  i = t
  describe(i)
  i.M()
}
func describe(i I) {
  fmt.Printf("(%v, %T)\n", i, i)
}
```
- Note that an interface value that holds a nil concrete value is itself non-nil.
- A `nil` interface value holds neither value nor concrete type.
#### Empty Interface
- The interface type that specifies zero methods is known as the empty interface: ``interface{}``
### Type Assertions
- A type assertion provides access to an interface value's underlying concrete value.
```go
t,ok := i.(T)
```
- This statement asserts that, interface `i` holds concrete Type `T` and assigns the underlying value to `t`, `ok` is `true` if assertion is true.
- If underlying value is not of type `T` Then this line triggers **panic** in case of `t:= i.(T)`.
```go
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
```
### `Stringer` Interface
- A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
```go
/*
type Stringer interface {
    String() string
}
*/
import "fmt"

type Person struct {
  Name string
  Age  int
}

func (p Person) String() string {
  return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
  a := Person{"Arthur Dent", 42}
  z := Person{"Zaphod Beeblebrox", 9001}
  fmt.Println(a, z)
}
```
---
## `Errors` Interface
Go programs express error state with error values.

The error type is a built-in interface similar to ``fmt.Stringer``. It containes `Error()` method.
```go
package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative numbers: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0{
		return 0, ErrNegativeSqrt(x)
	}
	z:=x
	for z*z-x > 0.0001 || x-z*z > 0.0001{
		z -= (z*z -x)/(2*x)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
/*
Output:
    1.4142278559705035 <nil>
    0 cannot Sqrt negative numbers: -2.000000
*/
```
---
## `Reader` Interface
- The io package specifies the io.Reader interface, which represents the read end of a stream of data.
- The io.Reader interface has a Read method:

```go
func (T) Read(b []byte) (n int, err error)
```
- Read populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an io.EOF error when the stream ends.

---
**Types like `maps`, `slices`, `channels`, `functions`, and `interfaces` contain pointers internally. When you pass them to a function, you’re copying the pointer, so both the original and the copy refer to the same underlying data.**