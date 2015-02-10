package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// you can get a pointer to a struct
	// you can acess a property
	q := &v
	q.X = 1e9
	fmt.Println(v)

	// struct literals are a thing
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		r  = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, r, v2, v3)

	// An array's length is part of its type, so arrays cannot be resized.
	// This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// arrays have a literal definition
	t := []int{2, 3, 5, 7, 11, 13}

	// rather than call it an array-access notation,
	fmt.Println("t ==", t)

	for i := 0; i < len(t); i++ {
		fmt.Printf("t[%d] == %d\n", i, t[i])
	}

	// like Python, you can have a slice
	fmt.Println("t[:2] == ", t[:2])
	fmt.Println("t[4:] == ", t[4:])

	// you can not have negative slices
	// fmt.Println("t[-1:] == ", t[-1:])
	// invalid slice index -1 (index must be non-negative)

	// the `make` fn creates a slice
	aa := make([]int, 5)
	printSlice("aa", aa)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	// because arrays are immutable, for a slice `s`,
	// cap(s) returns the capacity of that slice, which is static
	// len(s) returns the actual number of values that have been set

	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}

	// If the backing array of s is too small to fit all the given values a bigger array will be allocated.
	// The returned slice will point to the newly allocated array.
	var zz []int
	printSlice("zz", zz)

	// append works on nil slices.
	zz = append(zz, 0)
	printSlice("zz", zz)

	// the slice grows as needed.
	zz = append(zz, 1)
	printSlice("zz", zz)

	// we can add more than one element at a time.
	zz = append(zz, 2, 3, 4)
	printSlice("zz", zz)

	// through my learnings, i've discovered that
	// this is an array
	// var l [10]int
	// this is a slice
	// var m []int

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	ppow := make([]int, 10)
	for i := range ppow {
		ppow[i] = 1 << uint(i)
	}
	for _, value := range ppow {
		fmt.Printf("%d\n", value)
	}

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

type Vertex struct {
	X int
	Y int
}
