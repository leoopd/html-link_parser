package main

import (
	"fmt"
	"log"

	u "github.com/leoopd/html-link_parser/util"
)

func main() {
	// var list u.LinkedList

	htmlFile, err := u.HtmlReader("../html/ex2.html")
	if err != nil {
		log.Fatal(err)
	}
	parsedHtmlFile, err := u.HtmlParser(htmlFile)
	fmt.Printf("%#v", parsedHtmlFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// u.TagParser(parsedHtmlFile, &list)
	// fmt.Println(list.Head)
}