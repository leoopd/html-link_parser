package htmllinkparser

import (
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

func check(e error) error {
	if e != nil {
		return e
	}
	return nil
}

func HtmlReader(path string) (string, error) {
	text, err := ioutil.ReadFile(path)
	check(err)
	return string(text), nil
}

func TagParser(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {

		fmt.Printf("Pre: %#+v\n", n.Attr[0].Val)
		fmt.Printf("Pre: %#+v\n", n.FirstChild.Data)

		for j := n.FirstChild.NextSibling; j != nil; j = j.NextSibling {
			fmt.Printf("Post: %#+v\n", j.FirstChild.Data)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		TagParser(c)
	}
}

func HtmlParser(text string) (*html.Node, error) {

	doc, err := html.Parse(strings.NewReader(text))
	check(err)

	return doc, nil
}
