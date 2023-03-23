package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func runAndWaitHardLog() int {
	time.Sleep(time.Second * 1)
	return 10
}

func main() {

	for i := 0; i < 5; i++ {
		a := runAndWaitHardLog()
		log.Info("runAndWaitHardLog")
		log.Infof("result:%d", a)

	}
	fmt.Println("done")
}
