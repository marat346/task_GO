 package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
	)

func main() {
var s bytes.Buffer
  
fileName := "log.txt"
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