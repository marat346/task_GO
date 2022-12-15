package main

import (
	"fmt"
	"math/rand"
  "time"
)
func distance (x,y int) {
  x = 2 * x + 10
  fmt.Println(x)
  y = -3 * y - 5
  fmt.Println(y)
  }

func getPoint (n1,n2 int) (int,int){
  return rand.Intn(n1),rand.Intn(n2)
}
  
func main() {
    rand.Seed(time.Now().UnixNano())
    distance(getPoint(20,30))
  }