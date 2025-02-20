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
