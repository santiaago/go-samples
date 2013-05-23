package main

import (
	"github.com/santiaago/go-samples/pluralsight/helloworld/greeting"
	"fmt"
	"time"
)

func RenameToFrog(r greeting.Renamable){
	r.Rename("Frog")
}

func main() {

	// declare var	
	fmt.Println("\ndeclare a variable")
	var message string  = "hello go world"
	fmt.Println(message)
	
	// declare multiple vars
	fmt.Println("\ndeclare multiple variables")
	var a, b, c int = 1,2,3
	fmt.Println(message,a,b,c)
	
	// you don't need to specify the types:
	fmt.Println("\nYou do not need to specify types!")
	var message2 = "Hello go world"
	var d,e,f = 1, false,3
	fmt.Println(message2,d,e,f)
	
	// := to have declaration and initialization at one time
	fmt.Println("\ndeclaration and initialization")
	message3 := "Hola go mundo"
	g,h,i := 1,false,5
	fmt.Println(message3,g,h,i)

	// basic pointer example
	var greet *string = &message // pointer to a string
	fmt.Println(message,greet, *greet)
	
	// pointer assignment
	fmt.Println("\npointer assignment")
	*greet = "hi"
	fmt.Println(message,greet,*greet)
	
	// basic user types
	fmt.Println("\nbasic user types")
	var s = greeting.Salutation{"John", "Hello"}
	fmt.Println(s.Name,s.Greeting)
	var s2 = greeting.Salutation{Greeting: "Hello", Name:"Cook"}
	fmt.Println(s2.Name,s2.Greeting)
	s2.Name = "Aloha!"
	fmt.Println(s2.Name,s2.Greeting)

	// constants
	fmt.Println("\ncontants")
	fmt.Println(greeting.PI,greeting.Language)
	fmt.Println(greeting.A,greeting.B,greeting.C)

	// basic function declaration
	fmt.Println("\nbasic function declaration")
	var salutation = greeting.Salutation{"Bob", "Hello"}
	greeting.Greet(salutation, greeting.Print)
	greeting.Greet(salutation, greeting.PrintLine)
	greeting.Greet(salutation, greeting.CreatePrintFunction("!!!"))

	// branching
	fmt.Println("\nsbranching")	
	greeting.Greet_ex(salutation, greeting.PrintLine, true);
	greeting.Greet_ex(salutation, greeting.PrintLine, false);

	// switch type
	fmt.Println("\nswitch type")
	greeting.TypeSwitchTest("hello")
	greeting.TypeSwitchTest(salutation)
	greeting.TypeSwitchTest(2)
	greeting.TypeSwitchTest(2.9)

	// loops
	fmt.Println("\nnormal for loop")
	greeting.Greet_loop(salutation, greeting.PrintLine, true, 5)
	fmt.Println("\nwhile loop")
	greeting.Greet_whileloop(salutation, greeting.PrintLine, true, 5)	
	fmt.Println("\ninfinite loop with break")
	greeting.Greet_infloop(salutation, greeting.PrintLine, true, 3)	

	// ranges
	fmt.Println("\nranges")
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
	fmt.Println("\nSlicing slices:")
	slice = slice[1:2]// pos 1 included 2 excluded
	//slice = slice[:2] // first two
	//slice = slice[1:] // old after position 1
	greeting.Greet_slice(slice, greeting.CreatePrintFunction("?"), true)	

	// methods and interfaces
	fmt.Println("\nmethods and interfaces:")
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
	fmt.Println("\nInterfaces:")		
	// dereference the type or get the adress of this variable
	RenameToFrog(&salutations[0])
	salutations.Greet_slice_ex(greeting.CreatePrintFunction("?"), true)	
	
	// using write interface
	// use our salutations type as an argument to the Fprintf 
	fmt.Println("\nwriter interface")
	fmt.Fprintf(&salutations[0], "The count is %d", 10)
	fmt.Println("Fprintf does not print anythig, it sets format to &salutation[0]")
	salutations.Greet_slice_ex(greeting.CreatePrintFunction("?"), true)

	// goroutines 
	fmt.Println("\ngoroutines:")
	// run in his own thread and then continue execution
	go salutations.Greet_slice_ex(greeting.CreatePrintFunction("<C>"), true)
	salutations.Greet_slice_ex(greeting.CreatePrintFunction("?"), true)
	time.Sleep(100 * time.Millisecond)

	// channels solution without timers
	fmt.Println("\nchannels:")
	// create a channel
	done := make(chan bool)
	// create a concurrent anonimous func
	go func(){
		salutations.Greet_slice_ex(greeting.CreatePrintFunction("<C>"), true)
		// the value true is going in the channel done
		done <- true
	}()
	salutations.Greet_slice_ex(greeting.CreatePrintFunction("?"), true)
	// the following line is going to block until it can read a value from 'done' 
	<- done

	// buffered channel
	fmt.Println("\nbuffered channel:")
	// set 2 channels
	done2 := make(chan bool, 2)
	// create a concurrent anonimous func
	go func(){
		salutations.Greet_slice_ex(greeting.CreatePrintFunction("<C>"), true)
		// the value true is going in the channel done
		done2 <- true
		done2 <- true
		fmt.Println("Done!")
	}()
	salutations.Greet_slice_ex(greeting.CreatePrintFunction("?"), true)
	// the following line is going to block until it can read a value from 'done' 
	<- done2
}
