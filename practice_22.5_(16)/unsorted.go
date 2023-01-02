package main

import (
	"fmt"
	"math/rand"
  "time"
  
)
const n = 10
var value int

func main() {
  rand.Seed(time.Now().UnixNano())
  var a [n] int
  for i:= 0 ;i < n;i++ {
    a[i] = rand.Intn(10 * n)
  }
  fmt.Println(a)
  var value int

  fmt.Println("Введите число которое нужно найти:")
  fmt.Scan(&value)
  index := find(a,value)
  after := (n - 1) - index
  
  fmt.Printf("Индекс: %v\n", index)
  fmt.Println("Элементов после идекса:",index,"со значением :",value, "равно:",after )

}

  
func find (a[n] int, value int)(index int){
  index = 0
  for i:= 0;i < n; i++ {
    if a[i] == value {
      index = i
      break
    }
  }
  return
}