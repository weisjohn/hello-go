package main

import (
	"fmt"
	"strconv"
)

type RGB struct {
	r, g, b int64
}

func hexToRGBStruct(hex string) RGB {
	var r, err1 = strconv.ParseInt(hex[0:2], 16, 0)
	var g, err2 = strconv.ParseInt(hex[2:4], 16, 0)
	var b, err3 = strconv.ParseInt(hex[4:6], 16, 0)

	_ = err1
	_ = err2
	_ = err3

	rgb := RGB{r, g, b}
	return rgb
}

func hexToRGBInts(hex string) (int64, int64, int64) {
	var r, err1 = strconv.ParseInt(hex[0:2], 16, 0)
	var g, err2 = strconv.ParseInt(hex[2:4], 16, 0)
	var b, err3 = strconv.ParseInt(hex[4:6], 16, 0)

	_ = err1
	_ = err2
	_ = err3

	return r, g, b
}

func main() {
	var r, g, b = hexToRGBInts("ffffff")
	fmt.Println("ffffff", r, g, b)

	abc := hexToRGBStruct("ffffff")
	fmt.Println("ffffff", abc)
}
