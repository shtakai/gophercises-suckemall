package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com/", "The root url for building a sitemap")
	flag.Parse()

	fmt.Println(*urlFlag)

	/*
		1. get the web page
		2. parse all links on the page
		3. build proper url with our link
		4. filter out w/ different url
		5. find all pages w/ BFS
		6. print out XML
	*/
}
