package main

import (
	"fmt"
	"github.com/santiaago/go-samples/newmath"
)

type Hey struct{
	name string
	greeting string
}

func main() {
	fmt.Printf("hello, world. Sqrt(2) = %v\n", newmath.Sqrt(2))

	// declare var
	var message string  = "hello go world"
	fmt.Println(message)
	
	// declare multiple vars
	var a, b, c int = 1,2,3
	fmt.Println(message,a,b,c)
	
	// you don't need to specify the types:
	var message2 = "Hello go world"
	var d,e,f = 1, false,3
	fmt.Println(message2,d,e,f)
	
	// := to have declaration and initialization at one time
	message3 := "Hola go mundo"
	g,h,i := 1,false,5
	fmt.Println(message3,g,h,i)

	// basic pointer example
	var greeting *string = &message // pointer to a string
	fmt.Println(message,greeting, *greeting)
	
	// pointer assignment
	*greeting = "hi"
	fmt.Println(message,greeting,*greeting)
	
	// basic user types
	var s = Hey{"John", "Hello"}
	fmt.Println(s.name,s.greeting)
	var s2 = Hey{greeting: "Hello", name:"Cook"}
	fmt.Println(s2.name,s2.greeting)
	s2.name = "Aloha!"
	fmt.Println(s2.name,s2.greeting)
}



















