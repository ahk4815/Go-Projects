package htmlparser

import (
	"io"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

// This should return the []link
func parseHtmlTree(r io.Reader) (*html.Node, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// this should be a private method
func BuildLink(r io.Reader) []Link {
	n, _ := parseHtmlTree(r)
	links := make([]Link, 0)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			nodeLink := Link{Href: "", Text: ""}
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					nodeLink.Href = attr.Val
					break
				}
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				nodeLink.Text = c.Data
				break
			}
			links = append(links, nodeLink)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)
	return links
}
