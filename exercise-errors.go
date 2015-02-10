package main

import (
	"fmt"
	"math"
)

// this is how to implement an interface
type ErrNegativeSqrt struct {
	num float64
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e.num))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &ErrNegativeSqrt{x}
	}
	z, _z := 1.0, 0.0
	for math.Abs(z-_z) > 0.001 {
		_z = z
		z = z - (math.Pow(z, 2)-x)/2*z
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
