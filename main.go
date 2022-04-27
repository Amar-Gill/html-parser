package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func main() {

	file, err := os.Open("ex1.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		panic(err)
	}

	links := make([]Link, 0)

	var f func(*html.Node)
	f = func(n *html.Node) {
		var link Link
		hrefString := ""
		text := ""

		if n.Type == html.ElementNode && n.Data == "a" {
			hrefString = parseHref(n)
			text = parseText(n, text)
			link = Link{hrefString, text}
			links = append(links, link)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	fmt.Printf("%v", links)
}

func parseHref(n *html.Node) string {
	var v string
	for _, att := range n.Attr {
		if att.Key == "href" {
			v = att.Val
		}
	}
	return v
}

func parseText(n *html.Node, s string) string {

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			s += c.Data
		}
		s = parseText(c, s)
	}
	return s
}
