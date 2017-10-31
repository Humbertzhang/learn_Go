package main

import "fmt"

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	temp := 0
	rst := 1

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &temp)
		if temp%2 == 1 {
			rst *= temp
		}
	}
	fmt.Println("output:", rst)
}
