package htmllinkparser

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func HtmlReader(path string) (string, error) {
	text, err := ioutil.ReadFile(path)
	check(err)
	return string(text), nil
}

func LinkParser(text string) {

	doc, err := html.Parse(strings.NewReader(text))
	check(err)

	var f func(*html.Node)
	f = func(n *html.Node) {

		if n.Type == html.ElementNode && n.Data == "a" {

			fmt.Printf("Pre: %#+v\n", n.Attr[0].Val)
			fmt.Printf("Pre: %#+v\n", n.FirstChild.Data)

			for j := n.FirstChild.NextSibling; j != nil; j = j.NextSibling {
				fmt.Printf("Post: %#+v\n", j.FirstChild.Data)
			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
