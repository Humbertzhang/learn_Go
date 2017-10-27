/*
判断两个字符串是否只是字母顺序乱了
note:go中string的char为rune
*/
package main

import (
    "fmt"
    "reflect"
)

func main(){
    s1 := ""
    s2 := ""
    fmt.Printf("Print first string:")
    fmt.Scanf("%s", &s1)
    fmt.Printf("Print second string:")
    fmt.Scanf("%s", &s2)
    out := judge(s1, s2)
    fmt.Println(out)
}

func judge(s1 string, s2 string) bool {
    if s1 == s2 {
        return false
    }
    map1 := make(map[rune]int)
    map2 := make(map[rune]int)
    for _, char := range s1{
        map1[char]++
    }
    for _, char := range s2{
        map2[char]++
    }
    if reflect.DeepEqual(map1, map2) {
        return true
    }else{
        return false
    }
}
