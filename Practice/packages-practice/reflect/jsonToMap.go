package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	dataBytes := ReadFromJson("./packages-practice/reflect/data.json")

	var targetMap any

	json.Unmarshal(dataBytes, &targetMap)

	PrintJson(reflect.ValueOf(targetMap), 0)
}

func PrintJson(val reflect.Value, depth int) {
	//fmt.Println("The Kind is ", val.Kind())

	switch val.Kind() {

	case reflect.Map:

		for _, key := range val.MapKeys() {

			fmt.Println(strings.Repeat("\t", depth), key.String(), ":")
			PrintJson(val.MapIndex(key), depth+1)
		}
	case reflect.Slice:

		for i := 0; i < val.Len(); i++ {
			PrintJson(val.Index(i), depth+1)
		}
	default:
		fmt.Println(strings.Repeat("\t", depth), val.Interface())
	}
}

func ReadFromJson(fp string) []byte {
	data, err := os.ReadFile(fp)

	if err != nil {
		panic(err)
	}
	return data
}
