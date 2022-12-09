 package main

import (
	"fmt"
  "os"
	)

func main() {

  text := "access denial"
  
  f,err := os.Create("access.txt")
  if err := os.Chmod("access.txt",0444); err != nil {
    fmt.Println(err)
  }
  if err != nil{
    fmt.Println(err)
  }
  defer f.Close()

  f.WriteString(text)
  
}