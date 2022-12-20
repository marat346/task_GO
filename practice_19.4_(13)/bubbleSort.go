package main

import (
	"fmt"
)
  
func main() {
  
  var array = [10] int {3,1,9,4,7,6,5,8,2,0}
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
  fmt.Println(array)
 }