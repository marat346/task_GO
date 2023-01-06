package main

import "fmt"

func debugSlice(input[] int){
  fmt.Printf("data:%+v\n",input)
  fmt.Printf("len:%+v\n",len(input))
  fmt.Printf("cap:%+v\n",cap(input))

}

func main() {
  //copy()
s := [] int {1,2,3,4,5,6}
debugSlice(s)
  //2:4
d:= make([]int,2,2)
fmt.Println("Что получилось",d)
 
  // copy (куда d,откуда s[2:4] ,) 
  copy(d,s[2:4])

  d[0] = 10
  fmt.Println("S",s)
  fmt.Println("D",d)
}