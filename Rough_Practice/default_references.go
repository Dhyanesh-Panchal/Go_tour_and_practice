package main

import "fmt"

/*
Passed as Reference:

*/

func addElement(s []int, m map[string]int) {
	fmt.Println(s, len(s), cap(s))

	//
	s = append(s, 10)
	fmt.Printf("\n%p\n", &s)
	m["NewKey"] = 100
}

// Anamolous Behavior when append() recreates the underlying array when capacity surpasses, this new assignment is assigned to the local `s` and not to the mySlice2 passed in the func. To avoid this pass the pointer to the slice.

func modifySlice(s []int) {
	fmt.Printf("\n%p\n", &s)
	fmt.Println(s, len(s), cap(s))

	// Modify a value
	s[0] = 200

	// if the capacity limit reached, then the append will allocate new Underlying array and assign it to that s.
	s = append(s, 10)

	fmt.Printf("\n%p\n", &s)
	fmt.Println(s, len(s), cap(s))
}

func main() {
	//myArr := [...]int{1, 2, 3, 4}
	mySlice2 := make([]int, 5, 5)
	myMap := make(map[string]int)
	myMap["ExistingKey"] = 1
	fmt.Println(mySlice2, myMap)
	fmt.Printf("\n%p\n", &mySlice2)
	modifySlice(mySlice2)
	fmt.Println(mySlice2, mySlice2[:cap(mySlice2)], myMap)
}
