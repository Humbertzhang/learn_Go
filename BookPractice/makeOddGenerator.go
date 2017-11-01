package main

import "fmt"

//func() unit 是指闭包里面的那个函数，那个函数的返回值为unit
//闭包的里面函数只能叫func?
func makeOddGenerator() func() uint {
    i := uint(1)
    return func() (ret uint){
        ret = i
        i += 2
        return
    }
}

func main(){
    nextOdd := makeOddGenerator()
    for i:=0; i<5; i++ {
        fmt.Println(nextOdd())
    }
}
