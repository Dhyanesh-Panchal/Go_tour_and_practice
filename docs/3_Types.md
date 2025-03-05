# Basic Types in Go

```txt
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

The `int`, `uint`, and `uintptr` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.


### Boolean
- Go does not allow implicit conversion between `bool` and `integer` types (`0` and `1` cannot be used as `false` and `true`).
### Strings
- `strings` in Go are _not slices_. Instead, a string in Go is a **read-only byte array (immutable sequence of bytes)**.
  - A string is a data structure that holds:
    1. A **pointer** to an underlying byte array.
    2. A **length** that specifies how many bytes it contains.
- **Runes and bytes for Unicode**
```go

func main() {
	s := "你好" // string with Unicode characters
	fmt.Println(len(s))      // Output: 6 (Unexpected!)
	fmt.Println([]rune(s))   // Output: [20320 22909]
	fmt.Println(len([]rune(s))) // Output: 2 (Correct length)
}
```
  - "你好" consists of two Unicode characters, but UTF-8 encodes each as 3 bytes, so len(s) gives 6 instead of 2.
### `uintptr`
- `uintptr` is often misunderstood because **it’s not a pointer** but **just a numeric representation of an address.**
# Zero Values
Variables declared without an explicit initial value are given their zero value.

- `0` for numeric types,
- `false` for the boolean type, and
- `""` (the empty string) for strings.
- Pointers, slices, maps, channels, functions: `nil`
- The zero value of a struct is a struct where all fields are set to their respective zero values.


# Type conversion
The expression `T(v)` converts the value `v` to the type `T`.
- Unlike in C, in Go assignment between items of different type requires an explicit conversion, thre is **NO implicit Type conversion**.
- The conversion is allowed only if the value is representable in the target type. For instance, converting a very large value to a type that cannot represent such a large value would result in a **compile-time error**.
- For examples: [Ref](./Rough_Practice/Type_conversions.go)
- **Type Inference**: When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

## Range of Primitive Types
```txt
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE 754 32-bit floating-point numbers
float64     the set of all IEEE 754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```
