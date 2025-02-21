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

- `strings` in Go are _not slices_. Instead, a string in Go is a **read-only byte array (immutable sequence of bytes)**.
  - A string is a data structure that holds:
    1. A **pointer** to an underlying byte array.
    2. A **length** that specifies how many bytes it contains.

# Zero Values
Variables declared without an explicit initial value are given their zero value.

- `0` for numeric types,
- `false` for the boolean type, and
- `""` (the empty string) for strings.
- Pointers, slices, maps, channels, functions: `nil`
- The zero value of a struct is a struct where all fields are set to their respective zero values.


# Type conversion
The expression `T(v)` converts the value `v` to the type `T`.
- Unlike in C, in Go assignment between items of different type requires an explicit conversion.
- The conversion is allowed only if the value is representable in the target type. For instance, converting a very large value to a type that cannot represent such a large value would result in a **compile-time error**.

- **Type Inference**: When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

