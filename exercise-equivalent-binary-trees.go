package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

/*
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
*/

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Walker(t *tree.Tree) chan int {
	ch := make(chan int)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	one, two := Walker(t1), Walker(t2)
	for i := range one {
		j := <-two
		if i != j {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("same", Same(tree.New(1), tree.New(1)))
	fmt.Println("different", Same(tree.New(1), tree.New(2)))
}
