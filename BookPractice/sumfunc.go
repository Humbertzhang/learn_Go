package main

import "fmt"

func main() {
	s := []float64{1.0, 2.0, 3.0}
	ret := sum(s)
	fmt.Println(ret)
}

func sum(s []float64) float64 {
	ret := 0.0
	for i := 0; i < len(s); i++ {
		ret += s[i]
	}
	return ret
}
