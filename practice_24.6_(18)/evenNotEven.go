package main

import "fmt"

func evenNotEven (slice[] int)(one,second [] int)  {
 even := make([]int,0,0) 
 notEven := make([]int,0,0)
  
 for _,v := range slice{
    if v % 2 == 0 {
      even = append(even, v)
    } else {
      notEven = append(notEven, v)
    }
  }
  return even,notEven
}

func main() {
s := [] int {1,2,3,4,5,6,7,8,9,10}
fmt.Println(evenNotEven(s))
}