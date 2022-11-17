package main
 
import "fmt"
 
func main() {
  
  var discount float64
  var sum float64

  fmt.Println("Введите сумму товара :")
  fmt.Scan(&sum)

  fmt.Println("Введите скидку :")
  fmt.Scan(&discount)

  maxDiscount := sum * 0.3
  
  if maxDiscount >= discount && discount <= 2000 {
    fmt.Println("Вам положена скидка в размере" , discount, "рублей")
  } else {
    fmt.Println("Вам не положена скидка")
  }
}