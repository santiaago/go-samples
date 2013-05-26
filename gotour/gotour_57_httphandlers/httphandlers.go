package main

import (
	"net/http"
	"fmt"
)

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

func (h String) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, h)
}

func (h *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,h.Greeting, h.Punct, h.Who)
}

func main(){
	// your http.Handle calls here
	//http://localhost:4000/string
	http.Handle("/string", String("I'm a frayed knot."))
	http://localhost:4000/struct
	http.Handle("/struct", &Struct{"Hello",":","Gophers!"})
	http.ListenAndServe("localhost:4000",nil)
}

















