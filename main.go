package main

import (
  "fmt"
  "math"
 
func main() {
  
  var i uint8 = 10
  
  for i = 0 ; i != 254; i -- {
    fmt.Println(i)
  if i == 0 {
    fmt.Println("кола закончилась")
    continue
  }
  fmt.Printf("Осталось бутылок %d \n" , i)
}
  }