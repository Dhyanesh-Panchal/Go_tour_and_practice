# reflect

Reflection gives the ability to examine types at runtime. It also allows to examine, modify, and create variables, functions, and structs at runtime.

Reflection is built around three concepts:
- Types
- Values
- Kinds

## Types

The function `reflect.TypeOf(m)` takes in the variable and returns the object of `reflect.Type` type.
- ``reflect.Type`` is an interface with methods: [ref](https://pkg.go.dev/reflect#Type)
    - `Name()` returns the name of the type.  Some types, like a slice or a pointer, don’t have names and this method returns an empty string.
    - `Kind()` returns teh kind of the type. The kind is what the type is made of — a slice, a map, a pointer, a struct, an interface, a string, an array, a function, an int or some other primitive type.
    -  If you define a `struct` named `Foo`, the `kind` is `struct` and the `type` is `Foo`.
    - Always use the Kind of the type to determine which function to use on the `reflect.Type` object else it will panic.
Here is a good example examinar function which Examines the types recursively
```go
func examiner(t reflect.Type, depth int) {
	
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
	
	switch t.Kind() {
	
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		examiner(t.Elem(), depth+1)
		
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
			}
		}
	}
}
```


## Values

`reflect.ValueOf(var) ` creates a `reflect.Value` instance of the variable.
- If we want to modify the value, we can pass pointer of the variable.
- `Type()` can be used to get the type of the varibale
