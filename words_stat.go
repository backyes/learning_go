package main

import (
    "io"
    "strings"
    "os"
    "bufio"
    "fmt"
)

func main() {
    // open input file
    fi, err := os.Open("words_stat.go")
    if err != nil {
        panic(err)
    }
    // close fi on exit and check for its returned error
    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()

    // make a buffer to keep chunks that are read
    reader := bufio.NewReader(fi)
    stat_map := make(map[string]int64)
    for {
        buf, err := reader.ReadBytes('\n')
        if err != io.EOF {
          words := strings.Fields(string(buf[:]))
          for _, word := range words {
            if _, ok := stat_map[word]; ok {
              stat_map[word] ++
            } else {
              stat_map[word] = 1
            }
          }
        } else {
          break
        }
    }
    for key, value := range stat_map {
      fmt.Printf("%s:%d\n", key, value)
    }
}
