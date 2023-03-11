package main

import "fmt"

type Marat struct {
	Name string
	Age  int
}

func (a *Marat) PrintInfo(n string) string {
	return n
}

func main() {
	m := Marat{"Makar", 35}

	fmt.Println(m.PrintInfo("GOGO"))

}
