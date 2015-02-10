package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {

	t := []int{2, 3, 5, 7, 11, 13}
	printSlice("t", t)

	var u []int
	u = append(u, 1)
	printSlice("u", u)

}
