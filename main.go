package main

import (
	"fmt"
  "strings"
)

func main() {
  fmt.Println("Определение количества слов, начинающихся с большой буквы в строке:")
  
  
  line_1 := " Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software."
count:= 0
  
for len(line_1) > 0 {
    index := strings.Index(line_1," ")
    word := line_1[:index]
   if word == strings.Title(word) {
    count++
    }
  }

  fmt.Println("Строка содержит", count ,"слов с большой буквы.")
  }