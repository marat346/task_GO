 package main

import "fmt"

func main(){
  var a float64
  var b float64
  var c float64

  fmt.Println("Введите коэффициент А квадратного уравнения :")
  fmt.Scan(&a)

  fmt.Println("Введите коэффициент А квадратного уравнения :")
  fmt.Scan(&b)
  
  fmt.Println("Введите коэффициент А квадратного уравнения :")
  fmt.Scan(&c)

  pow(b ,b)
  fmt.Println(b)
  }