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

  fmt.Println("Введите коэффициент B квадратного уравнения :")
  fmt.Scan(&b)
  
  fmt.Println("Введите коэффициент C квадратного уравнения :")
  fmt.Scan(&c)
  
  f := math.Pow(b, 2)
  diskriminant := f -  4 * a * c
	fmt.Println(diskriminant)
  
  if diskriminant < 0 {
    fmt.Println("У уравнения нету корней,так как дискриминант отрицательный :",diskriminant)
  }
  if diskriminant > 0 {
    - b = b
    fmt.Println( b)
    d := (-b + math.Sqrt(f)) / (2 * a)
    fmt.Printf("%.1f",d)
  }
  if diskriminant == 0 {
    z := - b / 2 * a
    fmt.Printf("%.1f",z)
  }
}