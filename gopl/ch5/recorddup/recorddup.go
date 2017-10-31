/*
exerxise 5.2
*/
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
    tags := make(map[string]int)
    for tag, time := range recorddup(tags, doc) {
        if time >= 2 {
            fmt.Printf("%10s %10d\n", tag, time)
        }
    }
}

func recorddup(tags map[string]int, n *html.Node) map[string] int{
    if n.Type == html.ElementNode{
        tags[n.Data]++
    }
	if n.FirstChild != nil {
        tags = recorddup(tags, n.FirstChild)
    }
	if n.NextSibling != nil {
        tags = recorddup(tags, n.NextSibling)
    }

    return tags
}
