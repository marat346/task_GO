package main

import (
	"fmt"
)
  
func main() {
    
  var array = [10] int {3,1,9,4,7,6,5,8,2,0}
  var min int
  rsl := 0
  
  for i:= 0; i < len(array);i++ {
    
    for j:= 0; i < len(array) - i;i++{
      if array [j] > array [j + 1] {
      min = array [j + 1]
      array [j + 1] = array [j]
      array [j] = min
      rsl++
    }
    }
  }
  fmt.Println(rsl)
  fmt.Println(array)
 }