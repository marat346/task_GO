package main
 
import "fmt"
 
func main() {

 var numberJust int

  fmt.Println("Введите любое число до 10 :")
  fmt.Scan(&numberJust)
  
  for i:= 0;i <= numberJust;i++ {
    fmt.Println(i)
  }
}