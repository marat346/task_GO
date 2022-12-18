package main

import (
	"fmt"
)
  
func main(){
var oneArray = [5] int {3,2,5,1,4}
max := oneArray[0]
  
  for i:=0;i < len(oneArray) - 1;i++{
    if oneArray[i + 1] < max {
    max  = oneArray[i + 1]
  }
}
  fmt.Println("Вывод:",max )
}