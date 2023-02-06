package main

import (
	"flag"
	"fmt"
)

func main() {
	var str string
	var substr string

	flag.StringVar(&str, "str", "строка для поиска", "string to search in")
	flag.StringVar(&substr, "substr", "строка", "string to search for")
	flag.Parse()

	match := isSubstring(&substr, &str)

	if match {
		fmt.Printf("Строка <%s> является подстрокой <%s>\n", substr, str)
	} else {
		fmt.Printf("Строка <%s> не является подстрокой <%s>\n", substr, str)
	}
}

func isSubstring(substr *string, str *string) bool {
	substrRunes := []rune(*substr)
	strRunes := []rune(*str)

	j := 0
	for _, r := range strRunes {
		if r == substrRunes[j] {
			j++
			if j == len(substrRunes) {
				return true
			}
		} else {
			j = 0
		}
	}
	return false
}