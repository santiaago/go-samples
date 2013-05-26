package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct{
	r io.Reader
}

func rot13(b byte) byte{
	var first, second byte
	switch{
	case 'a' <= b && b <= 'z':
		first, second = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		first, second = 'A', 'Z'
	default:
		return b
	}
	return (b - first + 13)%(second - first + 1) + first
}

func (r rot13Reader) Read(t []byte)(n int, err error){
	n, err = r.r.Read(t)
	for i := 0; i<n; i++{
		t[i] = rot13(t[i])
	}
	return
}
func main(){
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}



















