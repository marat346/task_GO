package main

import "fmt"
 
func main() {
  
  fmt.Println("Введите месяц в году:")
  var month string
  fmt.Scan(&month)

  switch month {
    case "декабрь", "январь", "февраль":
    fmt.Println("зима")
    
    case "март", "апрель", "май" :
    fmt.Println("весна")
    
    case "июнь", "июль", "август":
    fmt.Println("лето")
  
    case "сентябрь", "октябрь", "ноябрь":
    fmt.Println("осень")
    default:
    fmt.Println("Неправильно ввели месяц")
  }
}