package main

import "fmt"

func main() {
	str := "Привет мир!"
	lenstr := len(str)
	fmt.Println(lenstr)

	strRune := []rune(str)
	fmt.Println(len(strRune))
}
