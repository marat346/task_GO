 package main

import "fmt"

func main(){
  var oneCoin int
  var secondCoin int
  var thirdCoin int
  var sum int

  fmt.Println("Введите номенал первой монеты :")
  fmt.Scan(&oneCoin)

  fmt.Println("Введите номенал второй монеты :")
  fmt.Scan(&secondCoin)

  fmt.Println("Введите номенал третий монеты :")
  fmt.Scan(&thirdCoin)

  fmt.Println("Введите сумму товара :")

  if oneCoin == sum  || (oneCoin + secondCoin) == sum{
    fmt.Println("Заберите товар")
  }
  if secondCoin == sum || (secondCoin + thirdCoin) == sum {
    fmt.Println("Заберите товар")
    
  } else {
    fmt.Println("Вход в программу")
  }
}