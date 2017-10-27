package main

import "fmt"

func main() {
  defer func() {
    //str即为recover()返回的，传递给panic的参数
    str := recover()
    fmt.Println(str)
  }()
  panic("PANIC")
}
