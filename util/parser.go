package htmllinkparser

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func HtmlParser(path string) {
	file, err := os.Open(path)
	check(err)
	r := bufio.NewReader(file)

	fmt.Println(r)
}
