package main

import "fmt"
import "math"

func main() {
	n := 0.0
	m := 0
	sum := 0.0
	fmt.Scanf("%f %d", &n, &m)
	for i := 0; i < m; i++ {
		sum += n
		n = math.Sqrt(float64(n))
	}
	fmt.Println("Sum:", sum)
}
