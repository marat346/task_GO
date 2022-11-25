package main

import (
	"fmt"
	"math"
)
func main() {
  var sum float64
  var percent float64
  var year float64 
  var income_month float64
  var rsl float64


  
  fmt.Println("Введите сумму которую хотите внести :")
  fmt.Scan(&sum)

  fmt.Println("Введите ежемесячный процент капитализации:")
  fmt.Scan(&percent)

  fmt.Println("Введите насколько лет планируете:")
  fmt.Scan(&year)

  for i := 0; i < 12 * int(year); i++ {
    income_month = (sum / 100) * percent 
    rsl += income_month
    fmt.Println("Доход за месяц :" ,income_month)
    sum += income_month
    sum = math.Round(sum * 100) / 100
    fmt.Println(rsl)
    fmt.Println(sum)
    }  
  fmt.Println(sum - rsl)
}