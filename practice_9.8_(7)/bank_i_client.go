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
  var bank float64
  var client float64
  var rsl_client float64

  fmt.Println("Введите сумму которую хотите внести :")
  fmt.Scan(&sum)

  fmt.Println("Введите ежемесячный процент капитализации:")
  fmt.Scan(&percent)

  fmt.Println("Введите насколько лет планируете:")
  fmt.Scan(&year)

  rsl_client = sum
  for i:= 0; i < 12 * int(year); i++ {
    income_month = (rsl_client / 100) * percent
    rsl_client += income_month
    }

  client = math.Round(rsl_client  * 100) / 100
  bank = client - rsl_client
  fmt.Println("====================================")
  fmt.Println("Для", sum, "рублей,со ставкой", percent, "% и ",year, "год клиент получит ", client, "и разницу банка составляет ", bank)
}