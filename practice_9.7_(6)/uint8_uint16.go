package main

import (
  "fmt"
  "math"
 )
func main() {
  
for i := 0; i <= math.MaxUint32;i++ {
     if i == math.MaxUint8 {
    fmt.Println("Приходится переполнений чисел типа uint8 :",i)
    continue
  }
 }
}