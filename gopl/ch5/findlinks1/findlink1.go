package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks : %v", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
    fmt.Println("---------------------------------")
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
    //exercise 5.1
    if n.FirstChild != nil {
        links = visit(links, n.FirstChild)
    }

    if n.NextSibling != nil {
        links = visit(links, n.NextSibling)
    }

    return links
}
