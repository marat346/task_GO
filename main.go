package main

import (
	"fmt"
  "strings"
	)

func main() {
	sentences := [4]string {"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
  chars := [5]rune {'H','E','L','П','М'}
  
  fmt.Println(parseTest(sentences,chars))
  }
  
func parseTest (sentences []string, chars []rune) int {
  for _,v := range sentences {
     index := strings.Index(v," ")
     word := v[index:]
     indexFindWord := strings.IndexRune(word,chars)
    return indexFindWord
     }
  }
   