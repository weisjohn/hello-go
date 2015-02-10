package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	if y < 0 {
		y = 1
	}
	if x < 0 {
		x = 1
	}

	var G uint8

	G = uint8(x ^ y)
	var A uint8 = 255

	A = A - G

	return color.RGBA{0, 0, G, A}
}

func main() {
	m := Image{400, 400}
	pic.ShowImage(m)
}
