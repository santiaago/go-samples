package main

import (
	"fmt"
	"code.google.com/p/go-tour/tree"
)

func walkImp(t *tree.Tree, ch chan int){
	if t.Left != nil{
		walkImp(t.Left,ch)
	}
	ch <- t.Value
	if t.Right != nil{
		walkImp(t.Right,ch)
	}
}
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	walkImp(t,ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	walk1,walk2 := make(chan int), make(chan int)
	
	go Walk(t1,walk1)
	go Walk(t2,walk2)

	for {
		v1, ok1 := <-walk1
		v2, ok2 := <-walk2
		if v1 != v2 || ok1 != ok2{
			return false
		}
		if !ok1{
			break
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

















