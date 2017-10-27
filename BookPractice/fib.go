package main

import "fmt"

func fib(n int) int{
    if n == 0{
        return 0
    }else if n == 1{
        return 1
    }else{
        return fib(n-1) + fib(n-2)
    }
}

func main(){
    ret := 0
    ret = fib(6)
    fmt.Println(ret)
}
