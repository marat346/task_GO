package main

import (
	"fmt"
	"math"
)
func main() {
  var sum float64
  var percent float64
  var year float64
  var finich float64
  
  fmt.Println("Введите сумму которую хотите внести :")
  fmt.Scan(&sum)

  fmt.Println("Введите ежемесячный процент капитализации:")
  fmt.Scan(&percent)

  fmt.Println("Введите насколько лет планируете:")
  fmt.Scan(&year)

  for i := 0; i <=12; i++{
    var rsl float64 = (sum / 100) * percent * year
    finich += rsl
    sum+= finich
    sum = math.Round(finich * 100) / 100
    fmt.Println(sum)
  }
  
}