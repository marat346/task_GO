package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Привет?Что?у вас"
	strSplitted := strings.Split(str, " ")
	strSplitted[1] = "vjkjrj"
	fmt.Println(len(strSplitted))
	fmt.Println(strSplitted)
}
