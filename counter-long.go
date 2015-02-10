package main

import (
	"fmt"
	"github.com/weisjohn/hello/counter"
	"time"
)

func main() {

	start := time.Now()

	summer := counter.Counter(100000000)
	sum := summer()

	fmt.Println("duration", time.Since(start))
	fmt.Println("sum", sum)

}
