package main

import "fmt"

func main(){
  var one int
  var second int
  var third int

  fmt.Println("Три числа.")
  fmt.Println("Введите первое число:")
  fmt.Scan(&one)

  fmt.Println("Введите второе число:")
  fmt.Scan(&second)

  fmt.Println("Введите третье число:")
  fmt.Scan(&third)

  if one > 5{
    if second > 5 {
      if third >5{
        fmt.Println("Нету чисел больше 5 ")
      }
    }
  }
  
}
