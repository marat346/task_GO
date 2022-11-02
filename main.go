package main

import "fmt"

func main() {
  var pointMoscow string = "Марат"
  var pointRyazan string = "Рязань"
  var middleSpeed int

  fmt.Println("Введите среднюю скорость движения из", pointMoscow, ":")
  fmt.Scan(&middleSpeed)
  end := middleSpeed * 2

  if end >= 200 {
      fmt.Println("Вы приехали в ", pointRyazan, "!!!")
    }
  
  fmt.Println("_________________________")
   fmt.Println("Итого вы проехали:",end ,"км")
}