package main

import "fmt"

func main() {

	var s string = "hello golang"
	r := []rune(s)

	fmt.Printf("String: %s, rune slice: %v\n", s, r)

}
