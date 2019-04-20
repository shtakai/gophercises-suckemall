package link

import "io"

// Link represents a tag
//  <a href="_href">_text</a>
type Link struct {
	Href string
	Text string
}

// Parse takes html document and returns
// slices of Link.
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
