package main

import (
	"fmt"
	"strings"
)

func main() {
  s:= "hello wordl"

  for strings.Contains(s," "){
    spaceIndex:= strings.Index(s," ")
    fmt.Println(s[:spaceIndex])
  }
  
}