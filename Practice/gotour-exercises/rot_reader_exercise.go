package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func translate_rot13(x byte)byte{
	switch {
		case (x>=65 && x <=77) || (x>=97 && x <=109):
			x=x+13
		case (x>77 && x <=90) || (x>109 && x <=122):
			x=x-13
	}
	return x
}

func (rot13 rot13Reader) Read (b []byte) (int,error){
	n,err := rot13.r.Read(b)
	if err != nil {
		return 0,err
	}
	for i:=0;i<n;i++{
		b[i] = translate_rot13(b[i])
	}
	return n,nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
