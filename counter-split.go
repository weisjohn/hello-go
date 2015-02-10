package main

import (
	"fmt"
	"github.com/weisjohn/hello/counter"
	"time"
)

func count(size int, c chan int) {
	summer := counter.Counter(size)
	c <- summer()
}

func main() {

	start := time.Now()

	c := make(chan int)

	size := 100000000
	routines := 16
	work_size := size / routines

	// fan out, distribute
	for i := 0; i < routines; i++ {
		go count(work_size, c)
	}

	// sync back up
	sum := 0
	for i := 0; i < routines; i++ {
		sum += <-c
	}

	fmt.Println("duration", time.Since(start))
	fmt.Println("sum", sum)

}
