package main

import (
	"os"
  "fmt"
  "io"
  )

func main() {
  f,err := os.Open("log.txt")
  if err != nil{
    fmt.Println("Ошибка открытия файла", err)
    return
  }
  defer f.Close()
  buf := make([] byte ,128)
  if _,err := io.ReadFull(f,buf); err != nil{
    fmt.Println("Не смогли прочитать последовательность байтов из файла", err)
    return
  }
  fmt.Printf("%s\n",buf)
}