package main

import (
  "fmt"
  "math"
 )
func main() {
  var a float64
  a = 0
  epsilon := 0.0001
  for i :=0; i <15;i++{
    a+= 0.1
    if math.Abs(a - 1) < epsilon{
      fmt.Println("a приблизительно равно 1")
    }
  }
  
  }