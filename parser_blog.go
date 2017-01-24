
// try to learning html parser 

package main

import (
    "io/ioutil"
    "log"
    "fmt"
    "net/url"
    "net/http"
)

func main() {
     u, err := url.Parse("http://bing.com/search?q=dotnet")
     if err != nil {
       log.Fatal(err)
     }
     u.Scheme = "https"
     u.Host = "google.com"
     q := u.Query()
     q.Set("q", "golang")
     u.RawQuery = q.Encode()
     fmt.Println(u)

     resp, err := http.Get(u.)
     if err != nil {
           log.Fatal("get %s failed", u)
         }
     defer resp.Body.Close()

     body, err := ioutil.ReadAll(resp.Body)
     if err != nil {
       log.Fatal(err)
     }
     fmt.Println(body)
}
