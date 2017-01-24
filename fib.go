
    package main

    import "fmt"

    func fib(value int64) int64 {
      if (value <= 0 || value == 1) {
        return 1
      } else {
        return fib(value-1) * value
      }
    }

    func main() {
      fmt.Printf("fib: %d ", fib(8))
    }
