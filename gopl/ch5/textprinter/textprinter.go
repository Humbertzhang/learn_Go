package main
//exercie 5.3
import (
    "fmt"
    "os"

    "golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "textprinter : %v", err)
		os.Exit(1)
	}
	for _, content := range textprinter(nil, doc) {
		fmt.Println(content)
	}
}

func textprinter(content []string, n *html.Node) []string {
    if n.Type == html.ElementNode {
        if n.FirstChild != nil {
            content = textprinter(content, n.FirstChild)
        }
        if n.NextSibling != nil {
            content = textprinter(content, n.NextSibling)
        }
        if n.FirstChild == nil {
            if n.Data != "script" && n.Data != "style" {
                for _, a := range n.Attr {
                    content = append(content, a.Val)
                }
            }
        }
    }else{
        if n.FirstChild != nil {
            content = textprinter(content, n.FirstChild)
        }
        if n.NextSibling != nil {
            content = textprinter(content, n.NextSibling)
        }

    }
    return content
}
