package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	// create our map
	var _map = make(map[string]int)
	a := strings.Fields(s)
	for _, value := range a {
		count, ok := _map[value]
		if !ok {
			count = 0
		}
		count += 1
		_map[value] = count
	}
	return _map
}

func main() {
	wc.Test(WordCount)
}
