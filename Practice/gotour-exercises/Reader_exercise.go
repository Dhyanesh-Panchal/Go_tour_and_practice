/*
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (reader MyReader) Read(b []byte) (int, error) {
	for indx, _ := range b {
		b[indx] = 'A'
	}

	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
