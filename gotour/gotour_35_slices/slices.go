package main

import (
	"code.google.com/p/go-tour/pic"
)

func Pic(dx, dy int)[][]uint8{
	a := make([][]uint8, dy)
	for i:= range a{
		a[i] = make([]uint8, dx)
	}
	for y , s := range a{
		for x := range s{
			s[x] = uint8(x * y)
		}
	}
	return a
}

func main(){
	pic.Show(Pic)
}

















