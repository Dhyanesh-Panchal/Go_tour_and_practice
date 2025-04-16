package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

func main() {
	dataBytes := ReadFromJson("./packages-practice/reflect/data.json")

	var targetMap any

	err := json.Unmarshal(dataBytes, &targetMap)
	if err != nil {
		log.Fatalf("can't unmarshal json: %v", err)
	}

	fmt.Println(targetMap)
	_, ok := targetMap.(map[string]interface{})
	fmt.Println(ok)

	//PrintJson(reflect.ValueOf(targetMap), 0)
}

func PrintJson(val reflect.Value, depth int) {
	//fmt.Println("The Kind is ", val.Kind(), "Depth is ", depth)

	switch val.Kind() {

	case reflect.Map:
		for _, key := range val.MapKeys() {

			fmt.Println(strings.Repeat("\t", depth), key.String(), ":")
			PrintJson(val.MapIndex(key), depth+1)
		}
	case reflect.Interface:
		if !val.IsZero() {
			PrintJson(val.Elem(), depth)
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			PrintJson(val.Index(i), depth+1)
		}

	default:
		fmt.Println(strings.Repeat("\t", depth), val.Elem())
	}
}

func ReadFromJson(fp string) []byte {
	data, err := os.ReadFile(fp)

	if err != nil {
		panic(err)
	}
	return data
}
