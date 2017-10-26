package main

import (
    "fmt"
    "os"
    )
func main(){
    for i,args := range os.Args[0:]{
        fmt.Printf("Arg[%d]:%s\n", i, args)
    }
}
