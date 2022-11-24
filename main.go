package main

import (
  "fmt"
  "math"
 )
func main() {
  
  fmt.Println("Введите х для которого необходимо рассчитать sin ")
  var x float64
  fmt.Scan(&x)

  fmt.Println("Введите в какую степень вводить :")
  var epsilon float64
  fmt.Println(&epsilon)

  epsilon = (1 / 10) * math.Pow(10,epsilon)
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