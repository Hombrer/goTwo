package main

import (
	"fmt"
	"strings"
)

func main() {

	var s string = "hello golang"
	r := []rune(s)

	fmt.Printf("String: %s, rune slice: %v\n", s, r)
	for _, value := range s {
		if value == 'g' {
			fmt.Println(strings.ToUpper(string(value)))
		}
	}

}
