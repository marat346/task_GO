package main

import (
	"fmt"
	"math/rand"
  "time"
  
)
const n = 10

func main() {
  rand.Seed(time.Now().UnixNano())
  var a [n] int
  for i:= 0 ;i < n;i++ {
    a[i] = rand.Intn(10*n)
  }
  // число есть в массиве
  value := a[n/2]
  index := find(a, value)
  fmt.Printf("Индекс: %v\n", index)
  //Число нету в массиве
  value = n * 20
  index = find(a,value)
  fmt.Printf("Индекс:%v\n", index)
}

  
func find (a[n] int, value int)(index int){
  index = -1
  for i:= 0;i < n;i++{
    if a[i] == value{
      index = i
      break
    }
  }
  return
}