package htmllinkparser

import (
	"log"
	"testing"

	u "github.com/leoopd/html-link_parser/util"
	"golang.org/x/net/html"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func TestHtmlReader(t *testing.T) {

	path := "./html/ex2.html"
	want := "<html>\n<head>\n<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css\">\n</head>\n<body>\n<h1>Social stuffs</h1>\n<div>\n<a href=\"https://www.twitter.com/joncalhoun\">\nCheck me out on twitter\n<i class=\"fa fa-twitter\" aria-hidden=\"true\"></i>\n</a>\n<a href=\"https://github.com/gophercises\">\nGophercises is on <strong>Github</strong>!\n</a>\n</div>\n</body>\n</html>"

	result, err := u.HtmlReader(path)
	if result != want || err != nil {
		t.Fatalf("HtmlReader(s) = %q, %v, want match for %#q, nil", result, err, want)
	}
}

func TestHtmlParser(t *testing.T) {

	readHtml := "<html>\n<head>\n<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css\">\n</head>\n<body>\n<h1>Social stuffs</h1>\n<div>\n<a href=\"https://www.twitter.com/joncalhoun\">\nCheck me out on twitter\n<i class=\"fa fa-twitter\" aria-hidden=\"true\"></i>\n</a>\n<a href=\"https://github.com/gophercises\">\nGophercises is on <strong>Github</strong>!\n</a>\n</div>\n</body>\n</html>"
	var want *html.Node = &html.Node{Parent: (*html.Node)(nil), FirstChild: (*html.Node)(nil), LastChild: (*html.Node)(nil), PrevSibling: (*html.Node)(nil), NextSibling: (*html.Node)(nil), Type: 0x2, DataAtom: 0x0, Data: "", Namespace: "", Attr: []html.Attribute(nil)}

	result, err := u.HtmlParser(readHtml)
	if result.Type != want.Type || err != nil {
		t.Fatalf("HtmlReader(s) = %v, %v, want match for %#v, nil", result, err, want)
	}
	if result.DataAtom != want.DataAtom || err != nil {
		t.Fatalf("HtmlReader(s) = %v, %v, want match for %#v, nil", result, err, want)
	}
	if result.Data != want.Data || err != nil {
		t.Fatalf("HtmlReader(s) = %v, %v, want match for %#v, nil", result, err, want)
	}
	if result.Namespace != want.Namespace || err != nil {
		t.Fatalf("HtmlReader(s) = %v, %v, want match for %#v, nil", result, err, want)
	}
}

func TestTagParser (t *testing.T) {
	
}
