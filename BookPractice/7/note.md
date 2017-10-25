# defer panic recover

defer 在函数停止后执行，即使出错或者Panic也会执行
***
panic 停止程序运行并抛出Run Timeerror
***
recover 处理run time error 
停止panic并返回传递给panic的参数
但是recover必须在defer中运行，因为panic会导致程序停止，若不用defer无法执行recover

package main

import "fmt"

func main() {
  defer func() {
    str := recover()
    fmt.Println(str)
  }()
  panic("PANIC")
}
