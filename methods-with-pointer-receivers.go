package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// scale is a method that receives a pointer
func (v *Vertex) PointerScale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println("ref does nothing to original:", v, v.Abs())
	v.PointerScale(5)
	fmt.Println("pointer changes values on original struct instance", v, v.Abs())
}
