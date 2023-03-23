package main

import "fmt"

func main() {
	mString := make(map[string]float64)
	mString["first"] = 1.0
	mString["second"] = 2.0
	fmt.Println(mString)

	// композитный литерал
	mString = map[string]float64{
		"third": 3.0,
		"forth": 4.0,
	}
	fmt.Println(mString, "композитный литерал")
}
