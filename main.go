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
	u.LinkParser(htmlFile)
}
