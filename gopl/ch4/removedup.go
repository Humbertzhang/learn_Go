package main

import "fmt"

func main(){
    s := []string{0:"qwer", 1:"qwert", 2:"qwert", 3:"qaz", 4:"qaz", 5:"wer", 6:"Humbert"}
    s2 := remove(s)
    fmt.Println(s2)
}

func remove(s []string) []string{
    retlen := len(s)
    for i := 1; i < retlen; i++ {
        if s[i] == s[i-1]{
            for j:= i; j < retlen - 1; j++ {
                s[j] = s[j+1]
            }
            retlen -= 1
        }
    }
    return s[:retlen]
}
