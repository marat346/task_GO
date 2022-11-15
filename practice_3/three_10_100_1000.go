package main

import "fmt"
 
func main() {
   
  one := 0
  second := 0
  third := 0
  
  for i := 0; i <= 1000; i ++ {
   fmt.Println(i)
    
  one ++
  if one == 10 {
      fmt.Println("Переменная ONE заполнилась на числе 10")
      continue
     } 
    
  second++
  if second == 100 {
      fmt.Println("Переменная Second заполнилась на 100")
      continue
    }

  third++ 
  if third == 999 {
    fmt.Println("Переменная Third заполнилась на 1000")
      continue
   } 
  }
}