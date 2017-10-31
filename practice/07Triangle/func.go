package main

import "fmt"

func main() {
	x := 0
	fmt.Scanf("%d", &x)
	for i := 0; i < x; i++ {
		var a [3]int
		flag := 1

		for j := 0; j < 3; j++ {
			fmt.Scanf("%d", &a[j])
		}
		for k := 0; k < 3; k++ {
			if (a[k] + a[(k+1)%3]) <= a[(k+2)%3] {
				flag = 0
			}
		}
		if flag == 0 {
			fmt.Println(a[0], " ", a[1], " ", a[2], "不可以组成三角形")
		} else {
			fmt.Println(a[0], " ", a[1], " ", a[2], "可以组成三角形")
		}
	}
}
