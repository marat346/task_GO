package main

import "fmt"
 
func main() {
   a :=0
   for numbers:= 0; numbers <= 999999;numbers++{
   
  numbers1 := numbers / 100000
  numbers2 := numbers % 100000 / 10000
  numbers3 := numbers % 10000 / 1000
  numbers4 := numbers % 1000 / 100
  numbers5 := numbers % 100 / 10
  numbers6 := numbers % 10
   
  if numbers1 == numbers6 && numbers2 == numbers5 && numbers3 == numbers4 {

      fmt.Println("Билетик Зеркальный")
      a++
    }
     fmt.Println(a)
   }
}