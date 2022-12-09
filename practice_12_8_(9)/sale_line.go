package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
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
  if scanner.Text() == exit {
    fmt.Println("Выход из  программы")
    break
  }
  file.WriteString(fmt.Sprintf("%v %s %v \n", number_line, time.Now().Format("2006-01-02 15:04:05"), scanner.Text()))
 }
}