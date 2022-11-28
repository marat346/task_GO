package main

import (
	"fmt"
	"strings"
)

func main() {
  s1:= "hello"
  s2:= "ll"

  i:= strings.Index(s1,s2)
  fmt.Println(i)
  
}