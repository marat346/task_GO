 package main

import (
	"fmt"
	"io/ioutil"
	"os"
	)

func main() {
  f,err := os.Open("sale_garage.txt")
  if err != nil {
    fmt.Println("Файл пуст!!!")
    return
  }
  defer f.Close()
  
  resultBytes , err := ioutil.ReadAll(f)
  if err != nil{
    panic(err)
  }
  fmt.Println("Сохраненный лог:")
  fmt.Println(string(resultBytes))
}