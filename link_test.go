package link

import (
	"log"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

var html1 = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>`

var nestedText = `<a href="/other-page">A link to <span>another page</span></a>`

func TestNodeText(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(nestedText))
	if err != nil {
		t.Error(err)
	}

	text := nodeText(doc)
	log.Println(text)
	if text != "A link to another page" {
		t.Fatalf("Failed to retreive node text")
	}
}

func TestFindLinkNodes(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(html1))
	if err != nil {
		t.Error(err)
	}
	linkNodes := findLinkNodes(doc)
	if len(linkNodes) < 1 {
		t.Fatalf("No <a></a> found in document")
	}
	for _, n := range linkNodes {
		if n.Data != "a" {
			t.Fatalf("Should only care about <a></a> element only")
		}
	}
}
