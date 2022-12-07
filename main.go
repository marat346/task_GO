package main

import (
	"fmt"
	"os"
  "strconv"
)

func main() {
  var line string
  var exit string = "exit"
  var number_line int = 0
  
  file,err := os.Create("sale_garage.txt")
  if err != nil {
    fmt.Println("Не удалось открыть файл")
    return
  }
  defer file.Close()
for {
  fmt.Println("Введите строку:")
  fmt.Scan(&line)
  number_line ++
  file.WriteString("Номер строки:\n")
  file.WriteString(strconv.Itoa(number_line))
  file.WriteString("\n")
  file.WriteString(line)
  file.WriteString("\n")
  file.WriteString(line)
  
  if line == exit {
    fmt.Println("Выход из  программы")
    break
  }
}
}