package main

import "fmt"
 
func main() {
  
  fmt.Println("Введите месяц в году:")
  var month string
  fmt.Scan(&month)

  switch month {
    case "декабрь", "январь", "февраль":
    fmt.Println("Время года — зима")
    
    case "март", "апрель", "май" :
    fmt.Println("Время года — весна")
    
    case "июнь", "июль", "август":
    fmt.Println("Время года — лето")
  
    case "сентябрь", "октябрь", "ноябрь":
    fmt.Println("Время года — осень")
    default:
    fmt.Println("Неправильно ввели месяц")
  }
}