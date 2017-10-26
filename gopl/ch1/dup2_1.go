package main

import (
    "bufio"
    "fmt"
    "os"
)

func main(){
    counts := make(map[string]int)
    files := os.Args[1:]

    if len(files) == 0{
        countLines(os.Stdin, counts, "Stdin")        //press ctrl + d to stop
    }else{
        for _, arg := range files{
            f, err := os.Open(arg)
            if err != nil{
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue        //继续读下一个文件
            }
            countLines(f, counts, arg)
            f.Close()
        }
    }

    for line, n := range counts {
        if n > 1 {      //n>1是因为需要找大于两行重复的内容
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func countLines(f *os.File, counts map[string]int, name string){
    //对f进行扫描，并把每一行放到input.Text()中，用map记录下来
    input := bufio.NewScanner(f)
    for input.Scan(){
        counts[input.Text()]++
		if(counts[input.Text()] == 2){
			fmt.Println("doc name:", name)
		}
    }
}

