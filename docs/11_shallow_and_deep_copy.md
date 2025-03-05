# Data Copy in golang
Here in golang there are 2 main types of copies:

## Shallow Copy
A shallow copy creates a new Instance of a variable but doesn't copy the underlying composite variables.
E.g.
```go
type Car struct {
 Brand  string
 Model  string
 Owners []string //change index 0 of 'Owners' from copy
}

func main() {
 car1 := Car{Brand: "Tesla", Model: "Model X", Owners: []string{"Alice", "Bob"}}
 car2 := car1 // Shallow copy
 //change index 0 of 'Owners' from copy
 car2.Owners[0] = "John"

 //the Owners from car1 was changed too
 fmt.Printf("car1 :%v , Car2: %v", car1, car2)
}
```
-  In the examples, the car2 doesn't recreate the underlying slice, instead it refers to the original underlying array.
## Deep Copy
A completely isolated copy of the original data.
Deep copy creates a new instance of a variable and copies all internal elements, no matter how ‘deep’ they are in the data structure.
### Ways to create a deep copy
#### Using serialization-deserialization.
We can use `encoding/json` or `encoding/gob` and use Marshal-Unmarshal or Encode-Decode resp. to create a copy.
E.g.
```go
func deepCopy(src,dest interface{})error{
	var buf bytes.Buffer
	
	// Encode into bytes
	err := gob.NewEncode(&buf).Encode(src)
	if err != nil{
		return err
    }
	
	// decode the bytes at new place. Create a proper copy.
	return gob.NewDecoder(&buf).Decode(dest)
}
```


