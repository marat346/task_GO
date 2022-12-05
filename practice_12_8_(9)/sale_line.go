package main

import (
	"fmt"
  "os"
	)

func main() {
  var line string
  fmt.Println("Введите строку:")
  fmt.Scan(&line)
  
  file,err := os.Create("sale_garage.txt")
  if err != nil {
    fmt.Println("Не смогли создать файл.", err)
    return
  }
  defer file.Close()
  
  file.WriteString(line)
  fmt.Println(file.Name())
}
  