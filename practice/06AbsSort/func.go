package main

import "fmt"
import "math"

func main(){
    n := 0;
    fmt.Scanf("%d", &n);
    /*
    var a [n] int;
    会报错:No non-constant array bound
    StackOverflow:
    You can't instantiate an array like that with a value calculated at runtime. Instead use make to initialize a slice with the desired length. It would look like this;
    所以应该用Make来造切片
    */
    a := make ([]float64, n);
    for i:=0; i < n; i++ {
        t := 0.0;
        fmt.Scanf("%f", &t)
        a[i] = t
    }
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            if math.Abs(a[i]) < math.Abs(a[j]) {
                temp := a[j]
                a[j] = a[i]
                a[i] = temp;
            }
        }
    }

    for k := 0; k < n; k++ {
        fmt.Print(a[k], " ")
    }
}
