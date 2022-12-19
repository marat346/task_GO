package main

import (
	"fmt"
)
  
func arrayRevers (input [10] int ) (output [10] int) {
    for i := 0;i < len(output) - 1 ;i++ {
    output[i] = input [9 - i]
      }
   return
  }
  
func main() {
    var numbers [10] int
    for i:= 0;i < len(numbers) - 1;i++ {
      fmt.Scan(&numbers[i + 1])
    }
    for _, end := range arrayRevers(numbers){
      fmt.Println(end)
    }
 }