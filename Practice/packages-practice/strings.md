# `strings`

## Types
`strings` package offers three major types:
#### Builder:
A Builder is used to efficiently build a string using `Builder.Write` methods. It minimizes memory copying.

It containes few general functions:
- All of these functions are a pointer-receiver to `Builder` type.
- `Cap() int`: returns the capacity of the Underlying byteslice.
- `Grow(n int)`: Increases capacity of b. It will increase the capacity only if _capacity-lenght < n_
- `Reset()` resets buffer to empty.
- `String() string` returns string.
- `Write(p []byte) (int, error)` Write appends the contents of p to b's buffer.
- `WriteByte(c byte) error` appends this byte to string, while `WriteRune(r rune) (int, error)` appends the UTF-8 encoding of Unicode code point r to b's buffer.
- `WriteString(s string) (int, error)` appends contents of s to b's buffer.

#### Reader
A Reader implements the `io.Reader`, `io.ReaderAt`, `io.ByteReader`, `io.ByteScanner`, `io.RuneReader`, `io.RuneScanner`, `io.Seeker`, and `io.WriterTo` interfaces by reading from a string.
- `Len() int` returns number of bytes of the unread portion of string.
- `Read(b []byte) (n int, err error)` Read reads up to len(`str`) bytes into `str`.
- `ReadAt(b []byte, off int64) (n int, err error)` Reads at the offset of `off`.
- `ReadByte() byte,error` reads and returns the next byte from the input or any error encountered.
- `Reset(s string)` will reset the reader to read s.
