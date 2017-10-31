//exercise 5.7
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)


func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		b := forEachNode(c, pre, post)
        if b == false {
            break
        }
	}

	if post != nil {
		post(n)
	}
}


var depth int

func startElement(n *html.Node) bool{
	if n.Type == html.CommentNode {
		fmt.Printf("%*s<%s>\n", depth * 2, "", n.Data)
	}else if n.Type == html.ElementNode {
		if n.Data == "a" && n.Attr != nil{
			fmt.Printf("%*s<%s>\n", depth*2, "", "<a href='...'>")
		}else {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		}
		depth++
	}
}

func endElement(n *html.Node) bool{
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
