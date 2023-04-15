package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-done:
				fmt.Println("exit")
				return

			default:
				fmt.Println("message")
				time.Sleep(time.Second)
			}
		}
	}()

	wg.Wait()

}
