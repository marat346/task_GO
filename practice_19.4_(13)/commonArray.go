package main

import (
	"fmt"
)
  
func main(){
var fourLength = [4] int {1,2,3,4}
var fiveLength = [5] int {5,6,7,8,9}
var commonArray = [9] int {0,0,0,0,0,0,0,0,0}
index := 0
  for i := 0; i < len(fourLength);i ++{
    commonArray[i] = fourLength[i]
    index++
    fmt.Println(index)
  }
  for i := 0;i < len(fiveLength);i++{
    fmt.Println(fiveLength[i])
    commonArray[index + i] = fiveLength[i]
  }
  
  fmt.Println(commonArray)
}