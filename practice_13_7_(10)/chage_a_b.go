package main

import "fmt"

  func chage (a,b *int) {
    *a = 20
    *b = 10
  }
   
func main (){
  a :=  10
  b := 20
  fmt.Println(a)
  fmt.Println(b)
  chage(&a,&b)
  fmt.Println(a)
  fmt.Println(b)
  
  }