package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int{
	f0 := 0
	f1 := 1
	return func() int{
		t := f0 + f1
		f0 = f1
		f1 = t
		return t
	}
}

func main(){
	f := fibonacci()
	for i := 0; i<10; i++{
		fmt.Println(f())
	}
}



















