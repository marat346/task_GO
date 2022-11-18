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

  if rsl <  math.MaxUint8 && rsl > 0 {
      fmt.Println("Лучше сохранить в тип Uint8")
    } else if rsl <  math.MaxUint16 && rsl > 0 {
       fmt.Println("Лучше сохранить в тип int16")
    } else if rsl < math.MaxInt32 && rsl < 0 {
    fmt.Println("Лучше сохранить в тип int32")
   } else if rsl > 0 {
    fmt.Println("Лучше сохранить в тип Uint32")
   } else {
    fmt.Println("Лучше сохранить в тип int8")
   }
  }