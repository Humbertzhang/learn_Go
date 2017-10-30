package main

import "fmt"
import "net/http"
import "encoding/json"
import "os"

func main(){
    getjson()
}

func getjson() (error){
    resp, err := http.Get("https://xkcd.com/571/info.0.json")

    if err != nil{
        return  err
    }
    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return fmt.Errorf("get fail : %s", resp.Status)
    }
    fmt.Println(resp)
    return nil
}
