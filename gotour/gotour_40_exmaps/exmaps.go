package main

import (
	"strings"
	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string] int{
	
	m := make(map[string] int)
	w := strings.Fields(s)
	for _,w := range w{
		m[w]++
	}
	return m
}
func main(){
	wc.Test(WordCount)
}



















