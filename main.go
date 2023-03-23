package main

import(
  "fmt"
  "log"
"os"
"time"
)

func runAdnWaitEasy() int{
  time.Sleep(time.Second * 1)
  return 10
}




func main(){
  file,err := os.OpenFile("logFile.txt",os.O_CREATE | os.O_WRONLY | os.O_APPEND,0666)
  if err != nil{
    log.Fatalf("error while opening a file:%v\n", err)
  }
defer file.Close()

log.SetOutput(file)

  for i:= 0;i < 5;i++{
    a := runAdnWaitEasy()
    log.Println("runAndWait finished...")
    log.Println("result", a)
  }
  fmt.Println("done")
}