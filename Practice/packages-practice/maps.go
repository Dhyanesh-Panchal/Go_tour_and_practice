package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	m := make(map[string]int)

	// Add some data
	m["Key1"] = 10
	m["Key2"] = 100

	// maps.All(m) -> returns an iterator on the map, order is random.
	for key, value := range maps.All(m) {
		fmt.Println(key, value)
	}

	// maps.Clone(m) -> returns a shallow copy of map.
	m1 := map[string]map[string]int{
		"M1": map[string]int{
			"Key1": 10,
			"Key2": 100,
		},
		"M2": map[string]int{
			"Key1": 100,
			"Key2": 1000,
		},
	}

	m1Copy := maps.Clone(m1)

	m1Copy["M1"]["Key1"] = 10000   // this will side effect
	m1Copy["M2"] = map[string]int{ // this will not
		"New": 1,
	}
	fmt.Println("M1: ", m1)
	fmt.Println("M1_copy: ", m1Copy)

	// maps.Collect() takes iterator containing 2 values and creates a map.
	//func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K]V
	s2 := []string{"zero", "one", "two", "three"}
	m2 := maps.Collect(slices.All(s2))
	fmt.Println("m1 is:", m2)

	// maps.Copy(dst,src) copies all the key-values from src map to dst map, existing keys are overridden.

	// maps.DeleteFunc(m, del_fn) takes in the map and a function func(k,v)bool, It deletes the entry for which this funciton returns true.
	fmt.Println(m["Unknown_Key"]) // prints zero value

	m["Unknown_Key"] = 500

	fmt.Println(m["Unknown_Key"])

	// lets delete this key from m1
	maps.DeleteFunc(m, func(k string, v int) bool {
		if k == "Unknown_Key" {
			return true
		}
		return false
	})

	fmt.Println(m["Unknown_Key"]) // zero value.

	// maps.Equal(m1,m2)bool, Equal reports whether two maps contain the same key/value pairs. Values Must be comparable, else...
	// maps.Equal(m1, m1_copy)  // This will give comple time error as the values are not 'comparable'.
	// reflect.DeepEqual()

	// maps.Insert(m,seq), Insert adds the key-value pairs from seq to m. seq is of type iter.Seq2[K, V] where K-> comparable, v-> any.

	// maps.Keys Returns the interator containing keys
	for k := range maps.Keys(m1) {
		fmt.Println(k)
	}

	// maps.Values
	for v := range maps.Values(m1) {
		fmt.Println(v)
	}

}
