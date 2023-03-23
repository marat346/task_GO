package main

import (
	"fmt"
	"strconv"
)

// Создание метода на другом Type(Типе) ,а не на струтуре
type MyInt int

func (mi MyInt) toString() string {
	return strconv.Itoa(int(mi))
}

func main() {
	var i MyInt = 10
	iStr := i.toString()
	fmt.Println(iStr)
}
