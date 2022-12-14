package main

import (
	"fmt"
	"log"

	u "github.com/leoopd/html-link_parser/util"
)

func main() {
	var list u.LinkedList

	htmlFile, err := u.HtmlReader("../html/ex5.html")
	if err != nil {
		log.Fatal(err)
	}
	parsedHtmlFile, err := u.HtmlParser(htmlFile)
	if err != nil {
		log.Fatal(err)
	}
	u.TagParser(parsedHtmlFile, &list)

	for i := list.Head; i != nil; i = i.Next {
		fmt.Println(i.Link)
	}
}
