package main

import "fmt"

func main() {
	// var mNil map[int]string
	// mNil[1] = "first"
	// неинициализированный map выдает ошибку на runtime

	//инициализированый map
	var mNil = make(map[int]string)
	mNil[1] = "first"
	fmt.Println(mNil)

}
