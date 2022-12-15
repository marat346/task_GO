package main

import (
	"fmt"
)

func main() {
  a := []string{"A", "B", "C"}
  s := a[len(a)-1] // C
  fmt.Println(s)
  }