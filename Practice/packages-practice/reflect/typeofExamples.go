package main

import (
	"fmt"
	"reflect"
	"strings"
)

type MyStruct struct {
	a int
	b string
	c []uint
}

func main() {
	fmt.Println("-----Reflect with primitive data types-----")
	var myuInt uint8 = 10

	fmt.Println("myuInt:", reflect.TypeOf(myuInt))

	fmt.Println("-----Reflect with slices-----")
	arr := [...]int{1, 2, 3}

	arrType := reflect.TypeOf(arr)

	fmt.Println(arrType)
	fmt.Println(arrType.Kind())
	fmt.Println(arrType.Elem())

	fmt.Println("-----Reflect with interface{}-----")

	var s = []any{1, "2", uint64(3)}

	typeS := reflect.TypeOf(s)
	fmt.Println("typeS Name:", typeS.Name(), " Kind:", typeS.Kind())

	valueS := reflect.ValueOf(s)

	fmt.Println(valueS.Kind())
	if valueS.Kind() == reflect.Slice {
		fmt.Println(valueS.Index(0))
	}

	//valueS

	fmt.Println("ValueS:", valueS.Type())

	fmt.Println("-----Reflect with struct-----")

	st := MyStruct{
		a: 1,
		b: "hello",
		c: []uint{1, 2, 3},
	}

	typeSt := reflect.TypeOf(st)
	totalFields := typeSt.NumField()
	fmt.Println("total fields:", totalFields)

	// Iterate over all fields to get field name
	for i := 0; i < totalFields; i++ {
		field := typeSt.Field(i)
		fmt.Println("Name: ", field.Name, "Type: ", field.Type)
		fmt.Println(field.Tag)
	}

	valueSt := reflect.ValueOf(st)
	fmt.Println("valueSt:", valueSt)
	for i := 0; i < valueSt.NumField(); i++ {
		fmt.Println(valueSt.Field(i))
	}

	fmt.Println("----- Reflect with maps -----")

	m := map[string]int{
		"mystr":            6,
		"myint":            34,
		"myboolean":        30,
		"myinterfaceslice": 10,
	}

	valueM := reflect.ValueOf(m)
	for _, i := range valueM.MapKeys() {
		fmt.Println(valueM.MapIndex(i))
	}

	fmt.Println(reflect.TypeOf(m).Elem())

}

func examineType(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type: ", t.Name(), "\t Kind: ", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		examineType(t.Elem(), depth+1)
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
