package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/shtakai/gophercises-suckemall/link"
)

/*
	Strategy
	1. get the web page
	2. parse all links on the page
	3. build proper url with our link
	4. filter out w/ different url
	5. find all pages w/ BFS
	6. print out XML
*/

func main() {
	urlFlag := flag.String("url", "https://gophercises.com/", "The root url for building a sitemap")
	flag.Parse()

	fmt.Println(*urlFlag)

	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	links, _ := link.Parse(resp.Body)

	for _, l := range links {
		fmt.Println(l)
	}
}
