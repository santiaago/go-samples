package main

import (
	"github.com/santiaago/go-samples/pluralsight/helloworld/greeting"
	"fmt"
)

func RenameToFrog(r greeting.Renamable){
	r.Rename("Frog")
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
	var salutation = greeting.Salutation{"Bob", "Hello"}
	greeting.Greet(salutation, greeting.Print)
	greeting.Greet(salutation, greeting.PrintLine)
	greeting.Greet(salutation, greeting.CreatePrintFunction("!!!"))

	// branching
	greeting.Greet_ex(salutation, greeting.PrintLine, true);
	greeting.Greet_ex(salutation, greeting.PrintLine, false);

	// switch type
	greeting.TypeSwitchTest("hello")
	greeting.TypeSwitchTest(salutation)
	greeting.TypeSwitchTest(2)
	greeting.TypeSwitchTest(2.9)

	// loops
	fmt.Println("normal for loop")
	greeting.Greet_loop(salutation, greeting.PrintLine, true, 5)
	fmt.Println("while loop")
	greeting.Greet_whileloop(salutation, greeting.PrintLine, true, 5)	
	fmt.Println("infinite loop with break")
	greeting.Greet_infloop(salutation, greeting.PrintLine, true, 3)	

	// ranges
	fmt.Println("ranges")
	slice := []greeting.Salutation{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What is up"},
	}

	// slices
	var sl []int
	sl = make([]int, 3, 10)
	sl[0] = 1
	sl[1] = 10
	sl[2] = 500
	//sl[3] = 23 this would fail at runtime
	// another way to init would be like this
	// sl2 := []int{1,10,500, 25}

	// you cannot do the following
	// slice[4] = "foo"
	// use append instead
	slice = append(slice,greeting.Salutation{"Frank","Hi"})
	// this will expand the slice
	slice = append(slice, slice...)
	// to remove elements
	slice = append(slice[:1],slice[2:]...)

	greeting.Greet_slice(slice, greeting.PrintLine,true)


	//slicing slices
	fmt.Println("Slicing slices:")
	slice = slice[1:2]// pos 1 included 2 excluded
	//slice = slice[:2] // first two
	//slice = slice[1:] // old after position 1
	greeting.Greet_slice(slice, greeting.CreatePrintFunction("?"), true)	

	// methods and interfaces
	fmt.Println("methods and interfaces:")
	salutations := greeting.Salutations{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What is up"},
	}
	salutations.Greet_slice_ex(greeting.CreatePrintFunction("?"), true)	
	fmt.Println("--")	
	salutations[0].Rename("John")
	salutations.Greet_slice_ex(greeting.CreatePrintFunction("?"), true)	
	
	// interfaces
	fmt.Println("Interfaces:")		
	// dereference the type or get the adress of this variable
	RenameToFrog(&salutations[0])
	salutations.Greet_slice_ex(greeting.CreatePrintFunction("?"), true)	
	
	// using write interface
	// use our salutations type as an argument to the Fprintf 
	fmt.Println("writer interface")
	fmt.Fprintf(&salutations[0], "The count is %d", 10)
	fmt.Println("Fprintf does not print anythig, it sets format to &salutation[0]")
	salutations.Greet_slice_ex(greeting.CreatePrintFunction("?"), true)	
	
}
