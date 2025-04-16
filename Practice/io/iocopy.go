package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("This is some string")

	if n, err := io.Copy(os.Stdout, reader); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Wrote ", n, "bytes to stdout")
	}
}
