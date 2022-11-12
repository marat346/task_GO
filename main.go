package main
 
import "fmt"
 
func main() {
    
  var ticket int
  var mirror bool = false
  
  fmt.Println("Введите номер билета :")
  fmt.Scan(&ticket)
  
  leftSide := ticket / 1000
  oneLeftSide := (ticket % 1000) / 100
  rightSide := (ticket % 100) / 10
  secondRightSide := ticket % 10

   if leftSide == secondRightSide && rightSide == oneLeftSide{
     mirror = true
     fmt.Println("Билетик Зеркальный")
  }
  
  if mirror {
    fmt.Println("Билетик Счастливый")
  }

  if (leftSide + oneLeftSide) != (rightSide + secondRightSide){
    fmt.Println("Билетик НЕСчастливый")
  }
}