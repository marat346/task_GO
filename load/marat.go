package main

import (
	"fmt"
	"os"
)

func main() {

	text := "Snowflakes are nice,\nSnowflakes are white.\nThey fall by day,\nThey fall at night."
	file, err := os.Create("poem.txt")
	if err != nil {
		fmt.Println("не смогли создать файл")
		return
	}
	defer file.Close()

	file.WriteString(text)
}
