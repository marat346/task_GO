package main

import (
	"fmt"
	"strconv"
	
)

func main() {
  line := "10a"
  
  i,err :=strconv.Atoi(line)
  if err != nil {
    fmt.Println("ошибка")
    return
  }
fmt.Println(i)
  
  }