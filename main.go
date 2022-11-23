package main

import (
  "fmt"
  "math"
 )
func main() {
 var epsilon float64
 epsilon = 0.0001
  
  fmt.Println("Введите х для которого необходимо рассчитать sin ")
  
  var x float64
  fmt.Scan(&x)
  
  result := 0.0
  prevResult := 0.0
  fact := 1
  k := 0 
  
  for {
    if k > 0 {
      fact *= 2 * k * (2*k + 1) 
    }
    result += math.Pow(-1,float64(k)) * math.Pow(x,float64(2 * k + 1)) / float64(fact)
    if math.Abs(result - prevResult ) < epsilon {
      fmt.Println(result)
      break
    }
    k++
    prevResult = result
  }
}