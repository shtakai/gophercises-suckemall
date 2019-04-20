package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represents a tag
//  <a href="_href">_text</a>
type Link struct {
	Href string
	Text string
}

// Parse takes html document and returns
// slices of Link.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc, "")
	return nil, nil
}

func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
