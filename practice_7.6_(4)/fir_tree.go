package main

import "fmt"
 
func main() {
  
  fmt.Println("Вывод ёлочки.")

  var height int 

  fmt.Println("Введите высоту ёлочки:")
  fmt.Scan(&height)

   for i := 1; i <= height; i++ {
        for j := 0; j < height-i; j++ {
            fmt.Print(" ")
        }
        for n := 0; n < i*2-1; n++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
  }