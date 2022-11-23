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

  if rsl > 0 {
    if rsl <= math.MaxUint8 {
      fmt.Println("Попробуйте лучше тип данных Uint8")
    } else if rsl <= math.MaxUint16 {
      fmt.Println("Попробуйте лучше тип данных Uint16")
    } else {
      fmt.Println("Попробуйте лучше тип данных Uint32")
    }
  } else if rsl >= math.MinInt8 {
    fmt.Println("Попробуйте лучше тип данных Int8")
  } else if rsl >= math.MinInt16 {
    fmt.Println("Попробуйте лучше тип данных Int16")
  } else {
    fmt.Println("Попробуйте лучше тип данных Int32")
  }
  }