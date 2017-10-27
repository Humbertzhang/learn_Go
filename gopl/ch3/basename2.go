package main

import "strings"
import "fmt"

func main(){
    s := ""
    fmt.Print("Input a string:")
    fmt.Scanf("%s", &s)
    s2 := basename(s)
    fmt.Printf("Basename:%s\n", s2)
}

func basename (s string) string {
    slash := strings.LastIndex(s, "/")
    s = s[slash+1:]
    if dot := strings.LastIndex(s, ".");dot >= 0{
        s = s[:dot]
    }
    return s
}
