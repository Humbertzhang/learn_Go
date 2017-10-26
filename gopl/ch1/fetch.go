package main

import (
    "fmt"
    "net/http"
    "os"
    "io/ioutil"
    "strings"
)

func main(){
    for _, url := range os.Args[1:]{
        if !strings.HasPrefix(url, "http://"){
            url = "http://" + url
        }
        resp, err := http.Get(url)
        if err != nil{
            fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
            os.Exit(1)
        }
        b, err := ioutil.ReadAll(resp.Body)
        resp.Body.Close()
        if err != nil{
            fmt.Fprintf(os.Stderr, "fetch : reading %v\n", err)
            os.Exit(1)
        }
        fmt.Println("%s", string(b))
        fmt.Println("HTTP Status Code:", resp.Status)
    }
}
