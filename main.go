package main

import(
  "fmt"
  "log"
"os"
"time"
)

func runAndWait() int {
  time.Sleep(time.Second * 1)
  return 10
}

func main(){

  file,err := os.OpenFile("test2.txt",os.O_CREATE | os.O_WRONLY | os.O_APPEND,0666)
  if err != nil{
    log.Fatalf("error while opening a file:%v\n", err)
  }

  log.SetOutput(file)
   for i:= 0;i < 5;i++{
     a:= runAndWait()
     log.Println("runAndWait finished...")
     log.Println("result", a)
}

fmt.Println("done")
}