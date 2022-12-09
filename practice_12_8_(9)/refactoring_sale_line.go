package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
  "io/ioutil"
  "bytes"
)

func main() {
  var exit string = "exit"
  var number_line int = 0
  var b bytes.Buffer
  
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("Введите строку:")
  
  for scanner.Scan() {
   number_line ++
    if scanner.Text() == exit {
        fmt.Println("Выход из  программы")
        break
  }
   b.WriteString(fmt.Sprintf("%v %s %v \n", number_line, time.Now().Format("2006-01-02 15:04:05"), scanner.Text()))
  }
  
  fileName := "sale_garage.txt"
   if err := ioutil.WriteFile(fileName,b.Bytes(),0666);err != nil {
       panic(err)
  }
  file, err := os.Open(fileName)
   if err != nil{
       panic(err)
  }
  defer file.Close()
  
  resultBytes , err := ioutil.ReadAll(file)
   if err != nil{
       panic(err)
  }
  fmt.Println("Сохраненный лог:")
  fmt.Println(string(resultBytes))
}