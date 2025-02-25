# Type Parameters
- Go functions can be written to work on multiple types using type parameters. The type parameters of a function appear between brackets, before the function's arguments.
```go
func Index[T comparable](s []T, x T) int
```
- This declaration means that `s` is a slice of any type `T` that fulfills the built-in constraint comparable. `x` is also a value of the same type.
- `comparable` is a useful constraint that makes it possible to use the `==` and `!=` operators on values of the type.

# Generic Types

```go
// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}
```
- `type any = interface{}`: `any` is an alias for `interface{}` and is equivalent to `interface{}` in all ways.

    