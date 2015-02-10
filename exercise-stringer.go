package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	var s []string
	for i := 0; i < 4; i++ {
		s = append(s, strconv.Itoa(int(ip[i])))
	}
	return strings.Join(s, ".")
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
