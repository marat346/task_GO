 package main

import "fmt"

func main(){
  var one int
  var second int
  var third int
  sum := 0

  fmt.Println("Три числа.Еще")
  fmt.Println("Введите первое число:")
  fmt.Scan(&one)

  fmt.Println("Введите второе число:")
  fmt.Scan(&second)

  fmt.Println("Введите третье число:")
  fmt.Scan(&third)

  if one >= 5 {
      sum++ 
    fmt.Println(sum)
  if second >= 5 {
     sum++
    fmt.Println(sum)
  }else if third >= 5 {
    sum++
    fmt.Println(sum)
  } else {
    fmt.Println("Чисел нету")
  }
}
  }