package main

import "fmt"

func use(s interface{}) func() string {
	s2 := s.(string)

	return func() string { return s2 }
}

func main() {
	data := []string{"Hello1", "Hello2", "Hello3", "Hello4", "Hello5", "Hello6", "Hello7"}

	for i, val := range data {
		//fmt.Println(i, &i, val, &val, &data[i])
		fmt.Println(i, &i)
		use(val)
	}
}
