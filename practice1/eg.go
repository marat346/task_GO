package main

import "fmt"

func main() {
  var one int
  var second int
  var third int
  
  fmt.Println("Баллы ЕГЭ.")
  fmt.Println("Введите результат первого экзамена:")
  fmt.Scan(&one)

  fmt.Println("Введите результат второго экзамена:")
  fmt.Scan(&second)

  fmt.Println("Введите результат третьего экзамена:")
  fmt.Scan(&third)

  fmt.Println("Сумма проходных баллов: 275")

  amount := one + second + third
  fmt.Println("Количество набранных баллов:",amount)
  if amount <= 275{
      fmt.Println("Вы не поступили.")
  } else {
    fmt.Println("Ура ,вы поступили!!!")
  }
}