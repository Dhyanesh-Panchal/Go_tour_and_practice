# Maps
```go
map[KeyType]ValueType
```
- Map types are reference types, like pointers or slices, and so the value of m above is nil; it doesn’t point to an initialized map. 
- A `nil` map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic.

## Initialization
1. Using `make()`:
```go
m = make(map[string]int)
```
- The make function allocates and initializes a hash map data structure and returns a map value that points to it.

2. Using `map` literals:
```go
m := map[string]int{
    "A": 3711,
    "B":   2138,
    "C": 1908,
}
```

**Key point**: `maps` in Go always use the **heap** for storage, regardless of whether they are initialized using a `map` literal or `make()`.

## Internals of Map
- At the core of every `map` is a **header** (defined in the Go runtime as the `hmap` struct) which contains metadata and pointers needed for map operations.
- The header itself is relatively small (typically a few dozen bytes) and may sometimes reside on the stack if it doesn’t escape, though its pointers always reference heap-allocated memory.
  - **count**: The number of key/value pairs currently stored.

  - **flags**:A set of flags used internally by the runtime (e.g., to mark special states).

  - **B**: A power-of-two exponent that determines the number of buckets (i.e., the number of buckets is 2^B).
  This is used to compute the bucket index from the hash value.

  - **hash0**: A random seed used to randomize the hash function, which helps protect against certain classes of attacks.
  - **buckets**: A pointer to an array of buckets. This array is allocated on the heap and grows as needed.
- Good Read for Internals of Map: [Ref](https://phati-sawant.medium.com/internals-of-map-in-golang-33db6e25b3f8)

## working with maps
```go
m["route"] = 66
i := m["route"]
```
If the requested key doesn’t exist, we get the value type’s zero value.

- built in `len(m)` function returns the number of items in map.
- built in `delete(m,"Key")` removes an entry from the map.
- The delete function doesn’t return anything, and will do nothing if the specified key doesn’t exist.

- To check if the key exists in the map, two value assignment can be used.
  - ``i, ok := m["key"]``, here the `ok` variable is `false` if the key doesn't exist.
- TO Iterate over the map, `range` can be used:
```go
for key, value := range m {
    fmt.Println("Key:", key, "Value:", value)
}
```
