package main

import (
	"fmt"
)

const line = 12
func main() {
  a:= [line] int {1,2,2,2,2,3,5,6,7,8,9,10}
  fmt.Println(a)
  var value int
  fmt.Println("Введите число:")
  fmt.Scan(&value)
  index := find(a,value)
  fmt.Println("Первое вхождение заданного числа в индексе:", index)
  }
  
func find (a[line] int, value int)(index int){
  index = 0
  min := 0
  max := line - 1
  for max >= min {
       middle := (max + min) / 2
        if a[middle] == value{
            index = middle
            break
        } else if a[middle] > value {
            max = middle - 1
        } else {
            min = middle + 1
        }
    }
    return
  }