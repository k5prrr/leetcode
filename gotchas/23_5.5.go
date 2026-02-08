package main

import (
	"fmt"
)

type SomeStruct struct {
	Value int
}

func CheckForNil(i interface{}) {
	if i == nil {
		fmt.Println("nil")
		return
	}

	fmt.Println("not nil")
}

func main() {
	var s *SomeStruct
	CheckForNil(s)
}
