package main

import (
  "flag"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
)

func Get(url string) {
  res, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  robots, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s", robots)
}

func main() {
  url := flag.String("url", "http://www.baidu.com", "URL link")
  flag.Parse()
  Get(*url)
}

