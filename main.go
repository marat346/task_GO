package main

import (
	"fmt"
)
  
  func arrayRevers (input [10] int ) (output [10] int) {
    for i := 0;i < output [len(output) -1] ;i++ {
    output[i] = output[len(output) - 1] - i
      fmt.Println(output[i])
     }
    return
  }
  
  func main() {
    arrayInput := [10] int {1,2,3,4,5,6,7,8,9,10}
    for _, conclusion := range arrayRevers(arrayInput){
      fmt.Println(conclusion)
    }
    }