 package main

import (
	"fmt"
  "time"
  "os"
	)

func main() {
  var line string
  var number_line int = 0
  var exit string = "exit"
  
  file,err := os.Create("sale_garage.txt")
  if err != nil {
    fmt.Println("Файл не удаеться открыть")
    return
  }
  defer file.Close()
for {
  fmt.Println("Введите строку:")
  fmt.Scan(&line)
  number_line ++
  file.WriteString(fmt.Sprintf("%v %s %v \n", line, number_line,time.Now().Format("2006-01-02 15:04:05")))
  if line == exit {
    fmt.Println("Выход из  программы")
    break
  }
}
}