package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// declared, uniniatilized map
var d map[string]Vertex

// map literal
var b = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// or alternatively
var c = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	d = make(map[string]Vertex)
	d["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(d["Bell Labs"])
	fmt.Println(b)
	fmt.Println(c)

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
