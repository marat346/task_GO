package main

import (
	"fmt"
  "strings"
	)

func main() {
	sentences := [4]string {"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}

for _,v := range sentences {
     index := strings.Index(v," ")
     word := v[index:]
     rRune := strings.IndexRune(word,'H')
     fmt.Println(rRune)
     }
}	