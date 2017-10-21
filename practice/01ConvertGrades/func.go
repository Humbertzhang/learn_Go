package main

import "fmt"

func main(){
    var n int;
    fmt.Print("Input n:")
    fmt.Scanf("%d", &n);
    for i:= 0; i < n; i++ {
        var in int;
        fmt.Print("Input the ", i , "th number:")
        fmt.Scanf("%d", &in);
        if in < 0 || in > 100 {
            fmt.Println("Score is error!");
        } else if in >= 90 {
            fmt.Println("A");
        } else if in >= 80 {
            fmt.Println("B");
        } else if in >= 70 {
            fmt.Println("C");
        } else if in >= 60 {
            fmt.Println("D");
        }else{
            fmt.Println("E")
        }
    }
    fmt.Println("Bye!")
}
