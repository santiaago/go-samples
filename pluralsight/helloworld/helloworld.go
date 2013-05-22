package main

import (
	"github.com/santiaago/go-samples/pluralsight/helloworld/greeting"
	"fmt"
)

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
	var greet *string = &message // pointer to a string
	fmt.Println(message,greet, *greet)
	
	// pointer assignment
	*greet = "hi"
	fmt.Println(message,greet,*greet)
	
	// basic user types
	var s = greeting.Salutation{"John", "Hello"}
	fmt.Println(s.Name,s.Greeting)
	var s2 = greeting.Salutation{Greeting: "Hello", Name:"Cook"}
	fmt.Println(s2.Name,s2.Greeting)
	s2.Name = "Aloha!"
	fmt.Println(s2.Name,s2.Greeting)

	// constants
	fmt.Println(greeting.PI,greeting.Language)
	fmt.Println(greeting.A,greeting.B,greeting.C)

	// basic function declaration
	var salutation = greeting.Salutation{"Tom", "Hello"}
	greeting.Greet(salutation, greeting.Print)
	greeting.Greet(salutation, greeting.PrintLine)
	greeting.Greet(salutation, greeting.CreatePrintFunction("!!!"))
	// branching
	greeting.Greet_ex(salutation, greeting.PrintLine, true);
	greeting.Greet_ex(salutation, greeting.PrintLine, false);
}
