package main

import "fmt"

func modify(s map[string]interface{}) {
	s["key2"] = append(s["key2"].([]map[string]interface{}), map[string]interface{}{"key3": 1})
}

func main() {
	s := map[string]interface{}{
		"key1": 1,
		"key2": []map[string]interface{}{},
	}

	modify(s)
	fmt.Println(s["key2"])
}
