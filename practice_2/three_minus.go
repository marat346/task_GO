 package main

import "fmt"

func main(){
  var one int
  var second int
  var third int

  fmt.Println("Введите первое число :")
  fmt.Scan(&one)

  fmt.Println("Введите второе число :")
  fmt.Scan(&second)

  fmt.Println("Введите третье число :")
  fmt.Scan(&third)

  if one > 0 || second > 0 || third > 0{
    fmt.Println("Одно число является положительным")
  }else{
    fmt.Println("Нету положительных чисел")
  }
}