package main

import (
	"fmt"
)
var global1 = 1
var global2 = 2
var global3 = 3 

func first (a int) int {
  return  a + global1 + global2
}

func second (a int) int {
  return a + global3 + global2
  }

func third (a int) int {
 return a + global3 +global1
}

func main() {
  fmt.Println(first(second(third(10))))
  }