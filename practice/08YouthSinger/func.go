package main

import "fmt"

func main() {
	x := 0
	fmt.Scanf("%d", &x)
	/*
	   声明变量
	   1 var x int
	   2 var x = 0
	   3 x := 0
	*/
	var max = 0.0
	var min = 0.0
	var cnt = 0.0
	for i := 0; i < x; i++ {
		temp := 0.0
		fmt.Scanf("%f", &temp)
		cnt += temp

		if i == 0 {
			max = temp
			min = temp
		} else {
			if temp > max {
				max = temp
			}
			if temp < min {
				min = temp
			}
		}
	}
	cnt = cnt - max - min
	ret := cnt / float64((x - 2))
	fmt.Println("平均分为:", ret)
}
