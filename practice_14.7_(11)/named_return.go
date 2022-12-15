package main

import (
	"fmt"
)

func sumNumber(a int) (c int) {
  c = a + 5
  return
}

func multiplicationNumber(b int) (s int) {
  s = b * 5
  return
}

func main() {
  fmt.Println(sumNumber(10))
  fmt.Println(multiplicationNumber(5))
  }