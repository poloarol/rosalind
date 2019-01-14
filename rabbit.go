package main

import (
    "fmt"
  )

func main(){
  rabbits := 33
  months := 2
  count := fib(rabbits, months)
  fmt.Println(count)
}

func fib(n int, k int) int {
  a, b := 1, 1

  for i := 0; i < n-1; i++ {
    a, b = b, b + (a*k)
  }

  return a
}
