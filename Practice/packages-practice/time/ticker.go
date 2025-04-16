package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		t := <-ticker.C

		fmt.Println("ticker ticked: ", t.Unix())
	}
}
