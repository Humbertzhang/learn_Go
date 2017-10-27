package main

import "fmt"

func main(){
    s := ""
    fmt.Scanf("%s", &s)
    s2 := makecomma(s)
    fmt.Printf("After Comma:%s\n", s2)
}

func makecomma(s string) string {
    n := len(s)
    if n <= 3 {
        return s
    }
    return makecomma(s[:n-3]) + "," + s[n-3:]
}
