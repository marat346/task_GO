package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var status int32 = 0

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&status, 1)
		}()
	}

	wg.Wait()
	fmt.Println(status)
}
