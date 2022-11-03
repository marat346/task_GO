package main

import "fmt"

func main() {
  var a int
  var b int 
  
 fmt.Println("Введиет число А :")
 fmt.Scan(&a)

fmt.Println("Введите второе число B :")
fmt.Scan(&b)

if a < b {
    fmt.Println("Число А меньше числа B")
  } else if a == b {
  fmt.Println("Число" , b ,"равно числу A")
}  else {
    fmt.Println("Число B меньше числа A")
  }
}