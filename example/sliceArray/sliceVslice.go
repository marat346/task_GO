package main

import "fmt"
func debugSlice(input[] int){
  fmt.Printf("data:%+v\n",input)
  fmt.Printf("len:%+v\n",len(input))
  fmt.Printf("cap:%+v\n",cap(input))

}

func main() {
 
  s := make([]int,0,0)
  s = append(s, 1)
  s = append(s, 2)
  // append(T[]type ,a....)
  s = append(s, 3,4,5,6)
  debugSlice(s)
  //testSlice = append(testSlice, s...) это слайс в слайсе
  // append долго работате(лучше создать 1 массив туда 2 массива)
  //cap(s) таким образом не выделяет памяти
  testSlice := make ([] int,0,cap(s))
  testSlice = append(testSlice, s...)
  debugSlice(testSlice)
}