// try to learn go with the background of html parser
// goquery 
package main

import (
  "flag"
  "log"
  "github.com/djimenez/iconv-go"
  "github.com/PuerkitoBio/goquery"
)

func html_parser(url string) {
     // Load the URL
     res, err := http.Get(url)
     if err != nil {
       log.Fatal("load http URL error")
         // handle error
     }
     defer res.Body.Close()

     // Convert the designated charset HTML to utf-8 encoded HTML.
     // `charset` being one of the charsets known by the iconv package.
     utfBody, err := iconv.NewReader(res.Body, charset, "utf-8")
     if err != nil {
       log.Fatal("convert to utf-8 eror")
         // handler error
     }

     // use utfBody using goquery
     doc, err := goquery.NewDocumentFromReader(utfBody)
     if err != nil {
       log.Fatal("create new document for goquery failed")
         // handler error
     }
     // use doc...
}

func GetCharset(url string) {
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

