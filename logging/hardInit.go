package main

import (
	"fmt"
	"time"

	lg "github.com/sirupsen/logrus"
)

func runAndWaitInit() int {
	time.Sleep(time.Millisecond * 1)
	return 10
}

// Это функция в GO ,которая запускается перед функций Main()
// func init() {
//Формат JSON
// log.SetFormatter(&log.JSONFormatter{})
//  }

func main() {

	for i := 0; i < 100; i++ {
		a := runAndWaitInit()
		lg.Info("runAndWaitHardLog")
		lg.Infof("result:%d", a)

		// Препруждение что мы приближаемяс к 90
		if i > 90 {
			lg.Warn("close to 100.....")
		}

		if i == 99 {
			lg.Fatal("reached 100")
		}
	}
	fmt.Println("done")
}
