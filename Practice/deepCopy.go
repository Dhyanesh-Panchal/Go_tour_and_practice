package main

import (
	"encoding/json"
	"fmt"
)

type MyStruct struct {
	Ch chan bool        `json:"-"`
	A  string           `json:"a"`
	B  *map[int]string  `json:"Key_b"`
	C  map[string][]int `json:"Key_C"` // "-" fields will marshal with zero value.
	d  bool             // Unexported fields will NOT marshal.
}

func main() {
	s := MyStruct{
		Ch: make(chan bool),
		A:  "ABab",
		B: &map[int]string{
			1: "one",
			2: "two",
			3: "three",
		},
		C: map[string][]int{
			"Key1": []int{1, 2, 3},
			"Key2": []int{2, 4, 6},
			"Key3": []int{3, 6, 9},
		},
	}
	fmt.Println(s)

	encoded, err := json.Marshal(&s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(encoded)

	var sDeepCopy MyStruct

	err = json.Unmarshal(encoded, &sDeepCopy)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("%T %+v", sDeepCopy, sDeepCopy)
	fmt.Println(sDeepCopy.A)
}
