package main

import (
	"log"

	u "github.com/leoopd/html-link_parser/util"
)

func main() {
	htmlFile, err := u.HtmlReader("./html/ex1.html")
	if err != nil {
		log.Fatal(err)
	}
	parsedHtmlFile, err := u.HtmlParser(htmlFile)
	if err != nil {
		log.Fatal(err)
	}
	link, text := u.TagParser(parsedHtmlFile)
	u.Insert(link, text)
}
