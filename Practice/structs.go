package main

import "fmt"

type Address struct {
	addressLine string
	city        string
	pin         int
}

type ConflictingAddress struct {
	addressLine string
}
type Person struct {
	Name    string
	Age     int
	Father  *Person
	Address // nested structure
	ConflictingAddress
}

type NewStruct struct {
	a int
	b []int
}

func (s NewStruct) Modify() {
	s.a = 1000
	s.b[2] = 100
}

func main() {
	father := Person{Name: "father_name", Age: 50}
	t := Person{"Dhyaensh", 20, &father, Address{"Addr_1", "Ahmedabad", 303030}, ConflictingAddress{"conflict_addr_1"}}

	fmt.Printf("%v", t)

	// Access the nested struct direct with fields
	fmt.Println("\nCity is:", t.city)
	// In case of Ambigous referece use the type name.
	fmt.Println("\nAddress is:", t.Address.addressLine)
	fmt.Println("\nAddress is:", t.ConflictingAddress.addressLine)
	fmt.Printf("\nFather: %v", t.Father)
	fmt.Printf("\n type: %T", t)

	fmt.Println("------------")

	s := NewStruct{
		a: 5,
		b: make([]int, 5),
	}

	fmt.Println(s)
	s.Modify()
	fmt.Println(s)

}
