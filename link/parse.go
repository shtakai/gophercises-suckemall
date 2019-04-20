package link

import (
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
	// 1. find <a> nodes in document
	// 2. for each link node
	//   2-a) build a link
	// 3. return slice of link
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = "TODO: Parse the text.."
	return ret
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
