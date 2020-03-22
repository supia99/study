package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(*tree.Tree, chan int)
	walker = func(t *tree.Tree, ch chan int) {
		if t == nil {
			return
		}
		walker(t.Left, ch)
		ch <- t.Value
		walker(t.Right, ch)
	}

	walker(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	Walk(t1, ch1)
	Walk(t2, ch2)
	for v1 := range ch1 {
		if v1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	k := make(chan int, 10)
	go Walk(tree.New(1), k)
	for v := range k {
		fmt.Println(v)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
