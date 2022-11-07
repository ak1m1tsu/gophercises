package htmlparser

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

// TODO: Write some tests

const testingTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Test</title>
</head>
<body>
    <section>
        <a href="text">First link</a>
        <a href="comment">Second link <!-- comment --></a>
        <a href="outer">Outer link<a href="inner">Inner link</a></a>
    </section>
</body>
</html>
`

func TestBuildLink(t *testing.T) {
	node := html.Node{
		Type: html.ElementNode,
		Data: "a",
		Attr: []html.Attribute{
			{
				Key: "href",
				Val: "https://github.com",
			},
		},
	}
	node.AppendChild(&html.Node{
		Type: html.TextNode,
		Data: "GitHub",
	})

	link := buildLink(&node)

	if link.Href != "https://github.com" && link.Text != "GitHub" {
		t.Errorf(
			"Expected href - 'https://github.com' and text - 'GitHub', but got href - %s and text - %s",
			link.Href,
			link.Text,
		)
	}
}

func TestGetText(t *testing.T) {
	textNode := html.Node{
		Type: html.TextNode,
		Data: "Text",
	}
	commentNode := html.Node{
		Type: html.CommentNode,
		Data: "Comment",
	}
	elementNode := html.Node{
		Type: html.ElementNode,
		Data: "a",
	}
	elementNode.AppendChild(&html.Node{Type: html.TextNode, Data: "Child text"})
	elementNode.AppendChild(&html.Node{Type: html.CommentNode, Data: "Child comment"})

	if getText(&textNode) != "Text" {
		t.Errorf("Expected text - 'Text', but got %s", getText(&textNode))
	}

	if getText(&commentNode) != "" {
		t.Errorf("Expected text - '', but got %s", getText(&commentNode))
	}

	if getText(&elementNode) != "Child text" {
		t.Errorf("Expected text - 'Child text', but got %s", getText(&elementNode))
	}
}

func TestLinkNodes(t *testing.T) {
	r := strings.NewReader(testingTemplate)
	doc, _ := html.Parse(r)

	nodes := linkNodes(doc)

	if len(nodes) != 4 {
		t.Errorf("Expected 4 nodes, but got %d", len(nodes))
	}

	if nodes[0].Attr[0].Key != "href" && nodes[0].Attr[0].Val != "text" {
		t.Errorf(
			"Expected key - 'href' and val - 'test', but got key - %s and val %s",
			nodes[0].Attr[0].Key,
			nodes[0].Attr[0].Val,
		)
	}

	if nodes[1].Attr[0].Key != "href" && nodes[1].Attr[0].Val != "comment" {
		t.Errorf(
			"Expected key - 'href' and val - 'test', but got key - %s and val %s",
			nodes[1].Attr[0].Key,
			nodes[1].Attr[0].Val,
		)
	}

	if nodes[2].Attr[0].Key != "href" && nodes[2].Attr[0].Val != "outer" {
		t.Errorf(
			"Expected key - 'href' and val - 'test', but got key - %s and val %s",
			nodes[2].Attr[0].Key,
			nodes[2].Attr[0].Val,
		)
	}

	if nodes[3].Attr[0].Key != "href" && nodes[3].Attr[0].Val != "inner" {
		t.Errorf(
			"Expected key - 'href' and val - 'test', but got key - %s and val %s",
			nodes[3].Attr[0].Key,
			nodes[3].Attr[0].Val,
		)
	}
}
