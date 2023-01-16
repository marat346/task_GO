package main

import "fmt"
func debugSlice(input[] int){
  fmt.Printf("data:%+v\n",input)
  fmt.Printf("len:%+v\n",len(input))
  fmt.Printf("cap:%+v\n",cap(input))

}

func main() {
s := [] int {1,2,3,4,5,6}
  //s 1,2,3,10,5,6
debugSlice(s)
 //d := s[1:4] - это SAP SLICE(ЕСЛИ МЕНЯЕШЬ В SAPSLICE,то ПОМЕНЯЕТСЯ В ИСХОДНОМ МАССИВЕ !!!!!!)

d := s[1:4] // это КОПИЯ исх.массива
  //2,3,10
  debugSlice(d)
//s
  //array <---------------------->
  //capacity 6

  //slice <----------------------->
  //capacity 6
  
  //newSlice <-------------->
  // capacity 5 (капатси среза лежащего исход. array минус индекс начало d:= s[1:3] )
e:= s [3:5]
  debugSlice(e)// пример

e[0] = 10
  fmt.Println("s",s)
  fmt.Println("d",d)
}