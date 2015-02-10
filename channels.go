package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	// channels must be made via `make()`
	c := make(chan int)

	// this is an example of divide and conquer through go routines
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	// this blocks until each routine is done
	x, y := <-c, <-c // receive from c

	// sum the two together
	fmt.Println(x, y, x+y)
}
