package main

import (
	"fmt"

	"github.com/weisjohn/crawler"
)

func main() {
	resources := crawler.Crawl("http://johnweis.com")

	for uri, hash := range resources {
		fmt.Printf("%s : %s\n", hash, uri)
	}
}
