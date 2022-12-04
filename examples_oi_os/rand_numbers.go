package main

import (
	"os"
  "fmt"
  "math/rand"
  "time"
  )

func main() {
 file,err := os.Create("log.txt")
  if err != nil{
    fmt.Println("не смогли создать файл",)
    return
  }
  defer file.Close()
  
  rand.Seed(time.Now().UnixNano())
  n := rand.Intn(101)
  fmt.Println("Введите число от 1 до 100")
  file.WriteString("Введите число от 1 до 100 \n")
  for{
    var answer int
    for {
      _,_ = fmt.Scan(&answer)
      file.WriteString(fmt.Sprintf("Введено число %d",answer))
      if answer < 1 || answer > 100{
        fmt.Println("Число должно быть в диапозоне от 1 до 100")
        file.WriteString("Число должно быть в диапозоне от 1 до 100\n")
      }else{
        break
      }
    }
    if answer == n{
      fmt.Println("Ура!!!Число угадано")
      file.WriteString("Ура!!!Число угадано\n")
      return
    }else if answer < n {
      fmt.Println("Загаданное число больше")
      file.WriteString("Загаданное число больше\n")
    }else{
      fmt.Println("Загаданное число меньше")
      file.WriteString("Загаданное число меньше\n")
    }
  }
}
  