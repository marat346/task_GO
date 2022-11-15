package main

import "fmt"
 
func main() {
 a:= 0
  for numbers:= 0; numbers <= 999999;numbers++{
    leftSide := numbers / 10000
     oneLeftSide := (numbers % 1000) / 100
     rightSide := (numbers % 100) / 10
     secondRightSide := numbers % 10
    
   if leftSide == secondRightSide && rightSide == oneLeftSide{
    a++
     fmt.Println("Билетик Зеркальный")
     fmt.Println(a)
  }
    
 }
   }