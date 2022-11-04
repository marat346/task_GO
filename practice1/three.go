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
     fmt.Println("Среди введённых чисел есть число больше 5.")
  } else if second > 5{
    fmt.Println("Среди введённых чисел есть число больше 5.")
  }else if third > 5 {
    fmt.Println("Среди введённых чисел есть число больше 5.")
  } else {
    fmt.Println("Чисел нету")
  }
  
}
