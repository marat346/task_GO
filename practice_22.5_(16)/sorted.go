package main

import (
	"fmt"
	"math/rand"
  "time"
  
)
const n = 12


func main() {
  rand.Seed(time.Now().UnixNano())
  var a [n] int
  for i:= 0 ;i < n;i++ {
    a[i] = 1 * i + rand.Intn(10)
  }
  fmt.Println(a)
  var value int
  fmt.Println("Введите число:")
  fmt.Scan(&value)
  index := find(a,value)
  fmt.Println("Первое вхождение заданного числа в индексе:", index)
  }
  
func find (a[n] int, value int)(index int){
  index = 0
  min := 0
  max := n - 1
  for max >= min {
    middle := (max + min) / 2
    if a[middle] == value{
      index = middle
      break
    } else if a[middle] > value{
      max = middle - 1
    } else {
      min = middle + 1
    }
    }
    return
  }