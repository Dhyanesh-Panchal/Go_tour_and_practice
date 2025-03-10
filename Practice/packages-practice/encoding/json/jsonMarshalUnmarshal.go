package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

type MyStruct struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Active      bool   `json:"active"`
	Scores      []int  `json:"scores"`
	Preferences struct {
		Theme         string `json:"theme"`
		Notifications bool   `json:"notifications"`
	} `json:"preferences,omitempty"`
	Tags      []string  `json:"tags,omitempty"` // omitempty will omit the field while Marshaling if empty.
	CreatedAt time.Time `json:"createdAt"`
}

func main() {
	data, err := os.ReadFile("./packages-practice/encoding/json/data_2.json")

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(data))

	var st MyStruct

	err = json.Unmarshal(data, &st)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", st)

	fmt.Println(reflect.TypeOf(st).FieldByName("CreatedAt"))
	fmt.Println(reflect.TypeOf(st).FieldByName("Preferences"))

	// Writing the Data
	newData := MyStruct{
		Id:     0,
		Name:   "New Person",
		Active: true,
		Scores: []int{1, 2, 3},
	}

	dataBytes, err := json.MarshalIndent(newData, "", "\t")

	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("./packages-practice/encoding/json/newWrittenData.json", dataBytes, 0644)
}
