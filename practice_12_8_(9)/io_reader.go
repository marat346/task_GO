 package main

import (
	"fmt"
  "os"
  "io"
	)

func main() {
  f,err := os.Open("sale_garage.txt")
  if err != nil {
    fmt.Println("Файл пуст!!!")
    return
  }
  defer f.Close()
  
  file, err := f.Stat()
	if err != nil {
		fmt.Println("Ошибка ,не смогли узнать размер файла")
	}
  fmt.Printf("Size: %d\n", file.Size())
  
  buf := make([] byte ,56)
  if _,err := io.ReadFull(f ,buf);err != nil{
  fmt.Println("Не смогли прочитать файл",err)
}
  fmt.Printf("%s\n",buf)
}