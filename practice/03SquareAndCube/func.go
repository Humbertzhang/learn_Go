package main

import "fmt"

func main() {
	var n int
	fmt.Println("输入N个数，输出其中偶数的平方和 与 奇数的立方和.")
	fmt.Print("Input n:")
	fmt.Scanf("%d", &n)
	fmt.Println("In")

	evensum := 0
	oldsum := 0

	for i := 0; i < n; i++ {
		var temp int
		fmt.Scanf("%d", &temp)
		if temp%2 == 0 {
			evensum += temp * temp
		} else {
			oldsum += temp * temp * temp
		}
	}
	fmt.Println("Evensum:", evensum, "Oldsum", oldsum)
}
