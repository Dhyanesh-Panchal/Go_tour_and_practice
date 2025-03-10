package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type MyStruct2 struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Active      bool   `json:"active"`
	Scores      []int  `json:"scores"`
	Preferences struct {
		Theme         string `json:"theme"`
		Notifications bool   `json:"notifications"`
	} `json:"preferences"`
	Tags []string `json:"tags"`
}

func main() {
	data, err := FetchData("./packages-practice/encoding/json/data_2.json")

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(data))

	var st MyStruct2

	decoder := json.NewDecoder(bytes.NewReader(data))

	err = json.Unmarshal(data, &st)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", st)

	fmt.Println(reflect.TypeOf(st).FieldByName("Name"))
	fmt.Println(reflect.ValueOf(st).FieldByName("Preferences"))

}
