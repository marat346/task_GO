package main

import (
	"fmt"
	"time"

	lg "github.com/sirupsen/logrus"
)

func runWaitHardMod() int {
	time.Sleep(time.Second * 1)
	return 10
}

func main() {
	for i := 0; i < 5; i++ {
		a := runWaitHardMod()
		lg.Info("runAndWaitHardLog")
		lg.Infof("result: %d", a)
	}

	fmt.Println("done")
}
