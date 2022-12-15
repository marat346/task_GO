package main

import "fmt"

func sum(a,b int){
  if a > b {
    c:= b
    b = a
    a = c
    }
  s:=0
  for i := a;i <= b;i++ {
   if i % 2 == 0 {
       s = s + i
      }
    } 
  fmt.Println(s)
  }

func main() {
  sum(8,12)
  sum(12,8)
  sum(7,12)
}