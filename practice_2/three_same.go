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

  if one == second  || second == third {
    fmt.Println("Введите другое число пожалуйста, числа совпадают")
  } else {
    fmt.Println("Вход в программу")
  }
}