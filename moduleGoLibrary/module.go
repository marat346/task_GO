package main

import (
	"fmt"
	sl "github.com/vpatsenko/skilllog"
	"log"
)

func main() {
	l, err := sl.NewLogger("log.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(l)

	l.Info("some info log")
	l.Error(fmt.Errorf("some error"))

	l.Warn("warn message")

	if err := l.Close(); err != nil {
		log.Fatalln(err)
	}
}
