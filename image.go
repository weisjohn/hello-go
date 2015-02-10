package main

import (
	"code.google.com/p/go-tour/pic"
	"math"
)

func one(x, y int) int {
	return x + y
}

func two(x, y int) int {
	return x - y
}

func three(x, y int) int {
	if y == 0 {
		y = 1
	}
	return x * y
}

func four(x, y int) int {
	if y == 0 {
		y = 1
	}
	return x ^ y
}

func five(x, y int) int {
	if y == 0 {
		y = 1
	}
	return int(math.Pow(float64(x+y), 2))
}

func six(x, y int) int {
	if y == 0 {
		y = 1
	}
	return int(math.Pow(float64(y-x), 2))
}

func seven(x, y int) int {
	if y == 0 {
		y = 1
	}
	return int(math.Pow(float64(y*x), 2))
}

func eight(x, y int) int {
	if y == 0 {
		y = 1
	}
	return (x + y) / 2
}

func nine(x, y int) int {
	if y == 0 {
		y = 1
	}
	return (x - y) / 2
}

func ten(x, y int) int {
	var a int
	if x > y {
		a = x
	} else {
		a = y
	}
	return int(a)
}

func eleven(x, y int) int {
	var a int
	if x < y {
		a = x
	} else {
		a = y
	}
	return int(a)
}

func twelve(x, y int) int {
	return int(math.Abs(float64(x - y)))
}

var sum_x, sum_y = 0, 0

func thirteen(x, y int) int {
	sum_x += x
	if x == 0 {
		x = 1
	}
	return sum_x / x
}

func fourteen(x, y int) int {
	sum_y += y
	if y == 0 {
		y = 1
	}
	return sum_y / y
}

func fifteen(x, y int) int {
	if x == 0 {
		x = 1
	}
	return int(math.Sin(float64(x+y)) * 1000.0)
}

func sixteen(x, y int) int {
	if x == 0 {
		x = 1
	}
	return int(math.Tan(float64(y/x)) * 10000.0)
}

func seventeen(x, y int) int {
	if x == 0 {
		x = 1
	}
	return int(math.Cos(float64(y*x)) * 10.0)
}

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	for i := range img {
		line := make([]uint8, dx)
		for j := range line {
			line[j] = uint8(fourteen(i, j))
		}
		img[i] = line
	}
	return img
}

func main() {
	pic.Show(Pic)
}
