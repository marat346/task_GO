package main

import (
	"fmt"
	"os"
  "strconv"
  "time"
  "bufio"
)

func main() {
 
  var exit string = "exit"
  var number_line int = 0
  
  file,err := os.Create("sale_garage.txt")
  if err != nil {
    fmt.Println("Не удалось открыть файл")
    return
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("Введите строку:")
for scanner.Scan() {
 number_line ++
  file.WriteString("Номер строки:\n")
  file.WriteString(strconv.Itoa(number_line))
  file.WriteString("\n")
  file.WriteString(scanner.Text())
  file.WriteString("Дата:\n")
  file.WriteString(time.Now().Format("2006-01-02 15:04:05"))
  file.WriteString("\n")
  
  if scanner.Text() == exit {
    fmt.Println("Выход из  программы")
    break
  }
}
}