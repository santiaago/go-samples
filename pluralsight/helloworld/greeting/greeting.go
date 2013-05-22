package greeting

import (
	"fmt"
)

type Salutation struct{
	Name string
	Greeting string
}

type Printer func(string)()

const (
	PI = 3.14
	Language = "Go"
)
const (
	A = iota // represents successive untyped integer constants
	B // assumes it is the same type as A
	C
)

func Print(s string){
	fmt.Print(s)
}

func PrintLine(s string){
	fmt.Println(s)
}

func CreatePrintFunction(custom string)Printer{
	return func(s string){
		fmt.Println(s + custom)
	}
}

func PrintCustom(s string, custom string){
	fmt.Println(s + custom)
}

func Greet(salutation Salutation, do Printer){
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting, "yoo")
	// use underscore to ignore variables
	do(message)
	do(alternate)
}

func Greet_ex(salutation Salutation, do Printer, isFormal bool){
	message, alternate := CreateMessage_ex(salutation.Name, salutation.Greeting)
	// use underscore to ignore variables
	if isFormal{
		do(message)
	}
	do(alternate)	
}

func CreateMessage(name string, greeting ...string) (message string,alternate string){
	// using variatic function
	// to get lenght of the greeting parameter use len() 
	fmt.Println(len(greeting))
	message = greeting[1]+ " " +name
	alternate = "HEY!" + name
	return
}

func CreateMessage_ex(name, greeting string) (message string,alternate string){
	message = greeting+ " " +name
	alternate = "HEY!" + name
	return
}
