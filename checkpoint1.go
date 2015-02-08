// everything that I know so far - 2015.02.07

// everything needs a package name
package main

// imports can be declared factor style
import (
	"fmt"
	"math"
)

// only import and declaration statements are allowed outside of a function body

// a variable declaration
var a int

// int is a shorthand for uint, which is either 32 or 64 bytes, depending on architecture

// int8 int16 int32 int64 all exist

// no float exists, you have to use float32 or float64

// bools are primatives
var b bool

// a variable declaration and assignment at once
// i'm not explicitly sure what this type c will be, but I think `int`
var c = 2

// doing this here
// d := 2
// is illegal - "non-declaration statement outside function body"

// variables can be declared in shorthand like
var e, f int

// also, you can declare and assign outside a function
var g, h int = 3, 4

// you can also declare and assign mixed types without specifying their type
var i, j, k = 1, true, "foo"

// structs exist - they are defined with their consituent properties
type Coord struct {
	lat, long float32
}

// declarations of structs can take place outside of functions
var dayton = Coord{39.759444, -84.191667}

// functions declare their name, args and types, and returns and types
// functions can have named return values, so that you have a "naked" return
// the return statement captures every name that is declared and returns those values in that order
// this seems nice, but i don't know how much I'll use it
func swap(x, y int) (a, b int) {
	b, a = x, y
	return
}

// ideally, you could do this, however it's illegal
// ./checkpoint1.go:xx: duplicate argument x
// ./checkpoint1.go:xx: duplicate argument y
// func swap(x, y int) (y, x int) {
// 	return
// }

// constants can be declared
const Pi = 3.14

// constants can also come from exports
// which tells me that imports are resolved before compilation (which makes sense)
// but implies that you can't have dynamic exports, as in, modify your export on the fly
// because you can't execute it
const Pi2 = math.Pi

// constants can not be structs, the next line is illegal
// const foo = Coord{39.759444, -84.191667}
// const initializer Coord literal is not a constant

// things can overflow and you get horrible results

// main is the main function, it doesn't return anything
func main() {

	// uninitialized variables are given a zero value
	fmt.Println("a =", a) // returns "a = 0"

	// when the var statement is like this, it's more like a `let`, I think?
	// this is called a short-assignment
	d := 2

	fmt.Println("d =", d) // returns "a = 0"

	fmt.Println("dayton is located at", dayton)

	// functions that return multiple values can be returned in one statement
	// i like this sugar
	y, z := swap(1, 9)
	fmt.Println("y is", y, "z is", z) // returns "y is 9, z is 1"
}
