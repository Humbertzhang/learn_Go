package main

import "fmt"

func main() {
	odd := 17
	even := 20
	judge(odd)
	judge(even)
}

func judge(num int) {
	if num%2 == 1 {
		fmt.Println(num, "is odd number.")
	} else {
		fmt.Println(num, "is even number.")
	}
}
