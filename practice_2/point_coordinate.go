 package main

import "fmt"

func main(){
  var x int
  var y int
  
  fmt.Println("Введите координату точки Х:")
  fmt.Scan(&x)

  fmt.Println("Введите координату точки У:")
  fmt.Scan(&y)

  if x > 0 && y > 0{
      fmt.Println("Точка предналежит координатной четверти 1")
              }
  if x < 0 && y > 0{
      fmt.Println("Точка предналежит координатной четверти 2")
  }
  if x < 0 && y < 0{
      fmt.Println("Точка предналежит координатной четверти 3")
  }
  if x > 0 && y < 0{
      fmt.Println("Точка предналежит координатной четверти 4")
  }
}