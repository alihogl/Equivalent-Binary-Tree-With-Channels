package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
   defer close(ch)
   var walk func(t *tree.Tree)

	walk = func( t *tree.Tree) {
		if t == nil {
		return
	}	
	
	walk(t.Left)
	walk(t.Right)
	ch <- t.Value
	}
walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool  {

	var ch1 chan int = make(chan int)
	
	var ch2 chan int = make(chan int)

	go Walk(t1,ch1)
	go Walk(t2,ch2)

var ch2Values []int

for v := range ch2 {
	ch2Values = append(ch2Values, v)
}
	counter:=0
	lengthTree:=0
	for  v1:= range ch1 {
		lengthTree++
		for _, v2:= range ch2Values {
			if(v1==v2) {
				counter++
				
			}
		}
	}

	return counter == lengthTree
}
// Creating a channel and sending to Walk function to walk the tree.
func crawl() {
	ch := make(chan int)

go Walk(tree.New(1), ch)

		for i:=0; i<10; i++ {
			fmt.Println(<-ch)
		}
}
// Sending trees to Same function to check whether trees same.
func check() {
	var t1 *tree.Tree =  tree.New(1)
	var t2 *tree.Tree =  tree.New(1)

	fmt.Print(Same(t1,t2))
}


func main() {
	crawl()
	check()
}
