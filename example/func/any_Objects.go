package main

import "fmt"

func anyObjects (a ... interface{}) {
  fmt.Println(a)
}

func main() {
 
  anyObjects(1,"222", [2] int {1,2})
}