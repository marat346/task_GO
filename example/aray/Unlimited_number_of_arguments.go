package main

import (
	"fmt"
)

func unlimitedArg (s... int) int{
  aa := 0
  for _,v := range s {
    aa = aa + v  
}
  return aa
}

func main() {
	fmt.Println(unlimitedArg(2,3))
  fmt.Println(unlimitedArg(2,3,4,5))
  fmt.Println(unlimitedArg(2,3,4,5,10,10,10,10,10))
  }