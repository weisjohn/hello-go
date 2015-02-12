package main

import (
	"fmt"
	"net/http"

	"github.com/weisjohn/htmlrefs"
)

func main() {
	resp, _ := http.Get("http://johnweis.com")
	refs := htmlrefs.All(resp.Body)

	for _, ref := range refs {
		fmt.Println(ref.Token, ":", ref.URI)
	}
}
