package main

import (
  "fmt"
  "math"
 )

func main() {
  var cycle_uint8 = 0
  var cycle_uint16 = 0
  
for i:= 0; i < math.MaxUint32; i++ {
  if i % math.MaxUint8 == 0 {
       cycle_uint8++
        continue
     }
  if i % math.MaxUint16 == 0 {
       cycle_uint16++
        continue
     }
  }
  fmt.Println(cycle_uint8)
  fmt.Println(cycle_uint16)
}