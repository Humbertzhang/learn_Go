package main

import "fmt"
import "bytes"

func main(){
    s := ""
    fmt.Scanf("%s", &s)
    s2 := makecomma(s)
    fmt.Printf("After Comma:%s\n", s2)
}

func makecomma(s string) string {
	n := len(s)
	if n <= 3{
		return s
	}
	var buf bytes.Buffer
    for i := 0; i<n; i++{
        buf.WriteByte(s[i])
        if (n - i) % 3 == 1 && i != n-1{
            buf.WriteByte(',')
        }
    }
    return buf.String()
}
