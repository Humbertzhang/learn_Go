package main

import "fmt"
import "crypto/sha256"

func main(){
    c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("w"))
    count := 0
    for i:=0; i<32; i++ {
        if c1[i] != c2[i]{
            count += 1;
        }
    }

    fmt.Println("不同bit数目:",count)
}
