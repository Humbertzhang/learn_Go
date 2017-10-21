package main

import "fmt"
import "math"

var PI float64 = 3.1415927;

func main(){
    var n int;
    fmt.Print("Input n:")
    fmt.Scanf("%d", &n);
    for i := 0; i<n; i++ {
        var temp float64;
        fmt.Scanf("%f", &temp);
        temp = getvolumn(temp);
        fmt.Printf("%.3f\n", temp);
    }
}

func getvolumn(r float64) float64 {
   var volumn = 4.0/3.0 * PI * math.Pow(r, 3);
   return volumn
}
