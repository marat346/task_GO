package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
  line := "10 "
  line = strings.Trim(line)
  i,err :=strconv.Atoi(line)
  if err != nil{
    fmt.Println("ошибка")
    return
  }
fmt.Println(i)
  
  }