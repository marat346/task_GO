package main

import (
	"fmt"
  "strings"
  "unicode"
	)

func main() {
	sentences := []string {"Hello two world", "Hello one Skillbox", "Привет зеленый Мир", "Привет Skillbox"}

    chars := []rune {'H', 'E', 'L', 'П', 'М'}

    fmt.Println(sentences, chars)
    result := parseTest(sentences[:], chars[:])
    fmt.Println(result)
  }
  
func parseTest(sentences []string, chars []rune) [][]int {
  var result [][]int
  for _,v := range sentences {
     index := strings.Index(v," ")
     word := v[index:]
     lowerWorld := strings.ToLower(word)
        tmp := []int{}
        for _, char := range chars {
            lowerChar := unicode.ToLower(char)
            tmp = append(tmp, strings.IndexRune(lowerWorld, lowerChar))
        }
        result = append(result, tmp)
    }
    return result
}    