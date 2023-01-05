package main

import "fmt"

func main() {
 
  var sliceExample[] int

newSliceExample := make([]int,0,10)
  
 fmt.Printf("%+v\n",sliceExample)
  fmt.Printf("%v\n",newSliceExample)

  newStringExample := make([]string,10,10)
  fmt.Printf("%v\n",newStringExample)
}