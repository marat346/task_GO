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
  fmt.Scan(&sum)
  
  if oneCoin == sum  || (oneCoin + secondCoin) == sum {
      fmt.Println("Заберите товар")
  } else if secondCoin == sum || (secondCoin + thirdCoin) == sum {
      fmt.Println("Заберите товар")
    } else if thirdCoin == sum || (thirdCoin + oneCoin) == sum {
      fmt.Println("Заберите товар")
  } else if (oneCoin + secondCoin + thirdCoin) >= sum {
      fmt.Println("Заберите товар")
  } else if (oneCoin + secondCoin + thirdCoin) < sum {
    fmt.Println("Не хватает денег")
  } else {
    fmt.Println("Сдачи нету")
   }
  }