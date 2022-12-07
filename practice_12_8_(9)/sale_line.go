 package main

import (
	"fmt"
  "time"
  "os"
	)

func main() {
  var line string
  var number_line string
  
  fmt.Println("Введите строку:")
  fmt.Scan(&line)
  
  fmt.Println("Введите номер строки:")
  fmt.Scan(&number_line)
  

  file,err := os.Create("sale_garage.txt")
  if err != nil{
    fmt.Println("Файл не удаеться открыть")
    return
  }
  defer file.Close()
  file.WriteString("Номер строки:\n")
  file.WriteString(number_line)
  file.WriteString(time.Now().Format("2006-01-02 15:04:05"))
  file.WriteString("\n")
  file.WriteString(line)

}