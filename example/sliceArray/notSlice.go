package main

import "fmt"

const size = 2
const newSize = 3

func newValue () int{
  return 7
}

func main() {
 data := [size] int {2,3}
 fmt.Println("data",data)

  var newData [newSize] int
  for i:=0 ; i < size ; i++{
    newData [i] = data[i]
  }
 newData[newSize -1] = newValue()
  fmt.Println("newData",newData)
}