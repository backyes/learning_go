// try to learn go with the background of html parser
// goquery 
package main

import (
  "flag"
  "log"
  "github.com/PuerkitoBio/goquery"
)

func Get(url string) {
  doc, err := goquery.NewDocument(url)
  if err != nil {
    log.Fatal("go query NewDocument failed")
  }

  dhead := doc.Find("head")
  dcharset := dhead.Find("meta[http-equiv]")
  charset, _ := dcharset.Attr("content")

  log.Println(charset)
}

func main() {
  url := flag.String("url", "http://www.baidu.com", "URL link")
  flag.Parse()
  Get(*url)
}

