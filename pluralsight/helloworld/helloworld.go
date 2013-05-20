package main

import (
	"fmt"
)

type Salutation struct{
	name string
	greeting string
}

const (
	PI = 3.14
	Language = "Go"
)
const (
	A = iota // represents successive untyped integer constants
	B // assumes it is the same type as A
	C
)

func Greet(salutation Salutation){
	fmt.Println(CreateMessage(salutation.name, salutation.greeting))
}

func CreateMessage(name, greeting string) string{
	return greeting+ " " + name
}

func main() {

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
	var s = Salutation{"John", "Hello"}
	fmt.Println(s.name,s.greeting)
	var s2 = Salutation{greeting: "Hello", name:"Cook"}
	fmt.Println(s2.name,s2.greeting)
	s2.name = "Aloha!"
	fmt.Println(s2.name,s2.greeting)

	// constants
	fmt.Println(PI,Language)
	fmt.Println(A,B,C)

	// basic function declaration
	var salutation = Salutation{"Tom", "Hello"}
	Greet(salutation)
}


















