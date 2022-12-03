package main

import (
	"os"
  "fmt"
  )

func main() {
  text := "Hello Марат"
  file,err := os.Create("hello.txt")
  if err != nil {
    fmt.Println("Не смогли создать файл.", err)
    return
  }
  defer file.Close()
  
  file.WriteString(text)
  fmt.Println(file.Name())
}
  