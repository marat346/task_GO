package main

import (
  "fmt"
  "math"
 )
func main() {
  var a int16
  var b int16

  fmt.Println("Минимальный тип данных.")
  fmt.Println("_______________________")
  fmt.Println("Введите первое число А :")
  fmt.Scan(&a)

  fmt.Println("Введите второе число В :")
  fmt.Scan(&b)
  rsl := int32(a) * int32(b)
  fmt.Println(rsl)

  if rsl > 0{
    
  }
  }