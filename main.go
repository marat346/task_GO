package main

import "fmt"

func main() {
  var Ivanov int 
  var Petrov int
  var Sidorov int

  fmt.Println("Какая зарплата у Иванова ?")
  fmt.Scan(&Ivanov)

  fmt.Println("Какая зарплата у Петрова?")
  fmt.Scan(&Petrov)

  fmt.Println("Какая зарплата у Сидорова?")
  fmt.Scan(&Sidorov)

  if Ivanov > Petrov{
     if Ivanov > Sidorov { 
       if Petrov > Sidorov {
         end := Ivanov - Sidorov
         fmt.Println("У Иванова самая высокая зарплата :",Ivanov,  "с разней " , end)
    } else {
       end1 := Ivanov - Petrov
         fmt.Println(" У Иванова самая высокая зарплата :",Ivanov,  "с разней " , end1)
    }
       } else {
      fmt.Println("У Сидорова самая высокая зарплата")
        }
    } else {
    if Petrov > Sidorov{
      fmt.Println("У Петрова самая высокая зарплата")
    } else {
      fmt.Println("У Сидорова самая высокая зарплата")
        }  
    }
    }
