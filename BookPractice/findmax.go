package main

import "fmt"

func main() {
	findmax(1, 2, 3, 4, 99, 50, 20)
}

func findmax(args ...int) {
	max := 0
	for i, num := range args {
		if i == 0 {
			max = num
		} else {
			if num > max {
				max = num
			}
		}
	}
	fmt.Println("Max:", max)
}
