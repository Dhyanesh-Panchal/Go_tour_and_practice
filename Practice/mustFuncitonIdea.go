package main

import (
	"fmt"
	"go_tour/utils"
	"os"
	"reflect"
	"strconv"
)

//
//// It is hard to check everywhere for errors in the code
//// Hence Must utility function is a simple Idea to Seprate the error handling logic from the main code.
//// NOTE: This should be use carefully as there can be scenarious where error need to handle in different way (NOT the common way done in the Must func)
//func Must[T any](intput T, err error) T {
//	if err != nil {
//		// Error handling logic
//		panic(err)
//	}
//	return intput
//}

func main() {
	var arg uint64
	if len(os.Args) > 1 {
		arg = uint64(utils.Must(strconv.Atoi(os.Args[1]))) // We can easily chain the outputs without worring about the error.
	}

	fmt.Println(arg, reflect.TypeOf(arg), len(os.Args), cap(os.Args))

}
