package main

import (
	"fmt"
  "strings"
)

func main() {
  fmt.Println("Определение количества слов, начинающихся с большой буквы в строке:")
  
  
  line_1 := " Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software."
  line_2 := strings.ToLower(line_1)
  count:= 0

  for len(line_1) > 0 {
    if line_1[:1] != line_2[:1] {
      count++ 
    }
    line_1 = line_1[1:]
    line_2 = line_2[1:]
  }
  fmt.Println("Строка содержит", count ,"слов с большой буквы.")
  }