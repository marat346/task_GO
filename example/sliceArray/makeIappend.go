package main

import "fmt"
func debugSlice(input[] int){
  fmt.Printf("data:%+v\n",input)
  fmt.Printf("len:%+v\n",len(input))
  fmt.Printf("cap:%+v\n",cap(input))

}

func main() {
 testSlice := make ([] int,0,0)
  //** append(testSlice(куда добавить),1 (что добавить)) "Эта функция возращает значение"
testSlice = append(testSlice,1)
  
for i:=0;i < 10;i++{
  testSlice = append(testSlice, i)
  debugSlice(testSlice)
}
  debugSlice(testSlice)
  
}