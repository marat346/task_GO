package main

import (
	"fmt"
  "strings"
  "strconv"
	)

func main() {
  s := "a10 10 20b 20 30c30 30 dd"

  for strings.Contains(s," ") {
       one_word := strings.Index(s," ")
       i,err := strconv.Atoi(s[:one_word])
       if err == nil {
           fmt.Println(i)
       }
       s = s[one_word + 1:]
     }
       i,err := strconv.Atoi(s)
     if err == nil {
           fmt.Println(i)
       }
}
  