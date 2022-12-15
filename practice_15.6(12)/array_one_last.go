package main

import (
	"fmt"
)
func last (arrayInput  [10] int) (arrayOutput [10] int) {
  for i := 0;i < arrayInput[len(arrayInput) -1];i++ {
    arrayInput[i] = arrayInput[len(arrayInput) -1]
  }
}


func main() {
 
  }