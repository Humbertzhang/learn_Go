package main

import "fmt"
import "net/http"

func main(){
    resp, err := getjson()
}

func getjson() (resp *Response, err error){
    resp, err := http.Get("https://xkcd.com/571/info.0.json")

    if err != nil{
        return nil, err
    }
    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("get fail : %s", resp.Status)
    }
    fmt.Println(resp)
}
