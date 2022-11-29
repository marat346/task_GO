package main

import (
	"fmt"
)

func main() {
 s:= "hello wordl"
for i := 0;i <12;i++{
  s =s[1:]
  fmt.Println(s)
}
  
} 