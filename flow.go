package main

import (
	"fmt"
	"math"
)

func main() {
	defer fmt.Println("done")

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// for is go's while
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// ifs are awesome
	if sum == 1024 {
		fmt.Println("Yatta!")
	}

	// inline if awesomeness
	if v := math.Pow(1, 2); v > 0 {
		fmt.Println(v)
	}
}
