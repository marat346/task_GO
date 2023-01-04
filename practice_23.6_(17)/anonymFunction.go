package main

import (
	"fmt"
)

const size = 10

func main() {
	var array = [size] int {3,1,9,4,7,6,5,8,2,0}
  
  anonym := func (f[size] int) [size] int { return bubbleSort(array)}
  fmt.Println(anonym(array))
}

func bubbleSort(array[size] int) [size] int {
  var min int 
  for i:= 0; i < len(array) - 1;i++ {
    for j:= 0; j < len(array)- 1 - i;j++ {
      if array [j] > array [j + 1] {
          min = array [j + 1]
          array [j + 1] = array [j]
          array [j] = min
        }
      }
    }
  return array
  }