package main

import (
    iconv "github.com/djimenez/iconv-go"
    "io/ioutil"
    "log"
    "net/http" 
)

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    resp, err := http.Get("http://www.baidu.com")
    if err != nil {
        log.Fatal(err)
    }
 
    defer resp.Body.Close()
    input, err := ioutil.ReadAll(resp.Body)
    out := make([]byte, len(input))
    out = out[:]
    iconv.Convert(input, out, "gb2312", "utf-8")
    ioutil.WriteFile("out.html", out, 0644)
}