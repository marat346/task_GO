 package main

import (
	"fmt"
	"math"
)

func main() {
  var a float64
  var b float64
  var c float64

  fmt.Println("Введите коэффициент А квадратного уравнения :")
  fmt.Scan(&a)

  fmt.Println("Введите коэффициент А квадратного уравнения :")
  fmt.Scan(&b)
  
  fmt.Println("Введите коэффициент А квадратного уравнения :")
  fmt.Scan(&c)
  
  f := math.Pow(b, 2)
  diskriminant := f - 4 * a * c 
	fmt.Println(diskriminant)
}