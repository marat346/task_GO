package main

import (
	"fmt"
	"time"

	lg "github.com/sirupsen/logrus"
)

func runLog() int {
	time.Sleep(time.Millisecond * 1)
	return 10
}

func init() {

	lg.SetFormatter(&lg.JSONFormatter{})
	lg.SetLevel(lg.WarnLevel)
}

func main() {
	for i := 0; i < 100; i++ {
		a := runLog()
		lg.Info("runAndWaitHardLog")
		lg.Infof("result:%d", a)

		if i > 90 {

			lg.WithFields(lg.Fields{
				"newLevel": i,
			}).Warn("close to 100....")
		}

		if i == 99 {
			lg.Fatal("reached 100")
		}
	}

	fmt.Println("done")
}
