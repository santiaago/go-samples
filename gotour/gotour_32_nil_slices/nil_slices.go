package main

import (
	"fmt"
)

func main(){
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil{
		fmt.Println("nil!")
	}
}

func printSlice(s string, x []int){
	fmt.Printf("%s len=%d cap=%d %v\n",s,len(x),cap(x),x)
}


















