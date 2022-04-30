package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/amar-gill/linkparser"
)

func main() {
	filename := flag.String("file", "ex1.html", "html file name")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	links, err := linkparser.ParseLinks(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(links)
}
