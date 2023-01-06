package main

import "fmt"
func debugSlice(input[] int){
  fmt.Printf("data:%+v\n",input)
  fmt.Printf("len:%+v\n",len(input))
  fmt.Printf("cap:%+v\n",cap(input))

}

func main() {
 
 s := make([]int,4,4)
  s[0] = 1
  s[1] = 2
  s[2] = 3
  s[3] = 4
  debugSlice(s)
//Рекомендуется пробегать по слайсам с помощью Range
  for i,v := range s{
    fmt.Println(i,v)
  }
  
}