package main

import (
	"fmt"
)

func adder() func(int) int{
	sum := 0
	return func(x int) int{
		sum += x
		return sum
	}
}

func main(){
	//The adder function returns a closure. 
	//Each closure is bound to its own sum variable.
	pos, neg := adder(), adder()
	for i := 0; i<10; i++{
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}



















