/*
As a simple way to play with functions and loops, implement the square root function using Newton's method.

In this case, Newton's method is to approximate Sqrt(x) by picking a starting point z and then repeating:


To begin with, just repeat that calculation 10 times and see how close you get to the answer for various values (1, 2, 3, ...).

Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small delta). See if that's more or fewer iterations. How close are you to the math.Sqrt?

Hint: to declare and initialize a floating point value, give it floating point syntax or use a conversion:
*/

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	// for math.Abs(z - x) > 0.1 {
	for i := 0; i < 1000000; i++ {
		z = z - (math.Pow(z, 2)-x)/2*z
	}
	return z
}

func SqrtAlt(x float64) float64 {
	z, _z := 1.0, 0.0
	for math.Abs(z-_z) > 0.001 {
		_z = z
		z = z - (math.Pow(z, 2)-x)/2*z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(SqrtAlt(2))
	fmt.Println(math.Sqrt(2))
}
