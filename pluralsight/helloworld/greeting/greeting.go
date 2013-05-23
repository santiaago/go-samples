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
	if prefix := GetPrefix(salutation.Name) ;isFormal{
		do(prefix + message)
	} else{
		do(alternate)
	}
}

func Greet_slice(salutation []Salutation, do Printer, isFormal bool){
	
	// not using index so use _ 
	for _, s:= range salutation{
		message, alternate := CreateMessage_ex(s.Name, s.Greeting)
		// use underscore to ignore variables
		if prefix := GetPrefix(s.Name) ;isFormal{
			do(prefix + message)
		} else{
			do(alternate)
		}
	}	
}

func Greet_loop(salutation Salutation, do Printer, isFormal bool, times int){
	message, alternate := CreateMessage_ex(salutation.Name, salutation.Greeting)
	for i:= 0; i< times; i++{
		// use underscore to ignore variables
		if prefix := GetPrefix(salutation.Name) ;isFormal{
			do(prefix + message)
		} else{
			do(alternate)
		}
	}
}

func Greet_whileloop(salutation Salutation, do Printer, isFormal bool, times int){
	message, alternate := CreateMessage_ex(salutation.Name, salutation.Greeting)
	i:= 0
	for i< times{
		// use underscore to ignore variables
		if prefix := GetPrefix(salutation.Name) ;isFormal{
			do(prefix + message)
		} else{
			do(alternate)
		}
		i++
	}
}

func Greet_infloop(salutation Salutation, do Printer, isFormal bool, times int){
	message, alternate := CreateMessage_ex(salutation.Name, salutation.Greeting)
	i := 0
	for {
		if i >= times{
			break;
		}
		// this skips the rest of the code and starts over a new iterations
		if i% 2 == 0 {
			i++
			continue
		}
		// use underscore to ignore variables
		if prefix := GetPrefix(salutation.Name) ;isFormal{
			do(prefix + message)
		} else{
			do(alternate)
		}
		i++
	}
}

func GetPrefix(name string)(prefix string){
	switch name{
	case "Bob": 
		prefix = "Mr "
		fallthrough
	case "Joe", "Amy": prefix = "Dr "
	case "Mary": prefix = "Mrs "
	default: prefix = "Duude "
	}
	return
}

func TypeSwitchTest(x interface{}){
	switch x.(type){
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Salutation:
		fmt.Println("Salutation")
	default:
		fmt.Println("unknown")
	}
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



