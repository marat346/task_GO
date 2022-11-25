package main

import (
  "fmt"
  "math"
 )
func main() {
  
  fmt.Println("Введите х для которого необходимо рассчитать sin ")
  var x float64
  fmt.Scan(&x)

  fmt.Println("Введите в какую степень возводить :")
  var level float64
  fmt.Scan(&level)

  var epsilon float64 = 1  / math.Pow(10,level)
  
  result := 1.0
  prevResult := 0.0
  fact := 1
  k := 1 
  
  for {
   	fact *= k
		result += math.Pow(x,float64(k)) / float64(fact)
		if math.Abs(result - prevResult) < epsilon {
			fmt.Print(result)
			break
    }
    k++
    prevResult = result
  }
}