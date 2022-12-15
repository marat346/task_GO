package main

import (
	"fmt"
)

func main() {
  
  var numbers [100] int
  var n = 10
  even := 0
  notEven := 0
  
  fmt.Println("Введите числа:")
  for i:= 0;i < n;i++ {
    fmt.Scan(&numbers[i])
  } 
  fmt.Println("Вывод:")
  
  for i := 0;i < n ;i++ {
    if numbers[i] % 2 == 0 {
        even++
    } else {
       notEven++
    } 
  }
  fmt.Printf("Четных чисел:%v\n",even)
  fmt.Printf("НЕчетных чисел:%v\n",notEven)
}