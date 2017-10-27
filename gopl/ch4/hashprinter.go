package main

import "fmt"
import "crypto/sha256"

func main(){
    fmt.Println("Start get input(input 'stop' to stop)")
    for true {
        ch := ""
        fmt.Scanf("%s", &ch)
        if ch == "stop" {
            break
        }
        hashch := sha256.Sum256([]byte(ch))
        fmt.Printf("%s -> %x\n", ch, hashch)
    }
}
