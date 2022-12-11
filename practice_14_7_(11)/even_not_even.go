package main

import "fmt"

func evenNotEven(a int) bool{
  if a % 2 == 0 {
      return true
   } else {
      return false
    }
  }
  func main() {
    fmt.Println(evenNotEven(3))
    fmt.Println(evenNotEven(4))
  }