package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	fmt.Printf("Input the id:")
	id := ""
	fmt.Scanf("%s", &id)

	for _, url := range os.Args[1:] {
		doc, err := getdoc(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't get html : %v", err)
			os.Exit(1)
		}
		node := ElementByID(doc, id)
		fmt.Println(node)
	}
}

func getdoc(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("get %s url error : %v", url, err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parse %s url's html error : %v", url, err)
	}
	return doc, nil
}

func ElementByID(n *html.Node, id string) string {
	if n.Type == html.ElementNode && n.Data == id {
		return n.Data
	}
	if n.FirstChild != nil {
		return ElementByID(n.FirstChild, id)
	}
	if n.NextSibling != nil {
		return ElementByID(n.NextSibling, id)
	}
    return "No that id"
}
