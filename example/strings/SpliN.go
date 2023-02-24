package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "a, b, c, d"
	strSplittedAfter := strings.SplitAfter(str, ",")
	fmt.Println(strSplittedAfter)
}
