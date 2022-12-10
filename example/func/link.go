package main

import "fmt"

func makeItZero(x int){
  x = 0
}

func main() {
  a:= 0
  makeItZero(a)
  fmt.Println(a)
}