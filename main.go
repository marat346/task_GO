package main

import "fmt"

func debugSlice(input[] int){
  fmt.Printf("data:%+v\n",input)
  fmt.Printf("len:%+v\n",len(input))
  fmt.Printf("cap:%+v\n",cap(input))
}
//первый сопоб
// этот метод нужен когда важен порядок в слайсе
func remove (slice[]int,index int)[]int {
  return append(slice[:index],slice[index+1:]...)
}
//второй способ
func anotherRemove(slice[]int,index int)[]int {
  slice[len(slice) - 1],slice[index] =
  slice[index] ,slice[len(slice) - 1]
  return slice[:len(slice) - 1]
}

func main() {
s := [] int {1,2,3,4,5,6}
debugSlice(s)
  // первый способ удаление из  массива
s= remove(s,2)
debugSlice(s)
 //второй способ удаление из массива
s = anotherRemove(s,3)
debugSlice(s)
}