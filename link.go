package link

import (
	"strings"

	"golang.org/x/net/html"
)

// Link is a representation of ancor tag (<a href=""></a>) in HTML document
type Link struct {
	Href string
	Text string
}

// ParseString take string of HTML document parse it and return
// all the links inside, return an error if it failed
func ParseString(source string) []Link {
	r := strings.NewReader(source)
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	linkNodes := findLinkNodes(doc)
	links := buildLinks(linkNodes)
	return links
}

func nodeText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += nodeText(c)
	}
	return text
}

func buildLinks(nodes []*html.Node) []Link {
	var links []Link
	for _, n := range nodes {
		var link Link
		link.Text = nodeText(n)
		for _, a := range n.Attr {
			if a.Key == "href" {
				link.Href = a.Val
				links = append(links, link)
				break
			}
		}
	}
	return links
}

func findLinkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var links []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, findLinkNodes(c)...)
	}
	return links
}
