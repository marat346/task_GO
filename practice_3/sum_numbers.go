package main
 
import "fmt"
 
func main() {

  var a  int
  var b int
  
  fmt.Println("Введите число А :")
  fmt.Scan(&a)

  fmt.Println("Введите число B :")
  fmt.Scan(&b)

  fmt.Println("___________________________")
  
  for i:= a;i <= b;i++ {
    a++
    fmt.Println(a)
  }
}