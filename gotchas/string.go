package main

import (
	"fmt"
)

func main() {
	v := "Hello🌍"
	for i, c := range v {
		fmt.Printf("%d of %d '%s'\n", i, c, string(c))
	}

	emoji := []rune("cool😎")
	for _, ch := range emoji {
		fmt.Printf("%d of '%s'\n", ch, string(ch))
	}
}
