package main

import (
	"fmt"
)

var m [2][2]string

func main() {

	m[0][0] = "v"
	m[0][1] = "v"
	m[1][0] = "v"
	m[1][1] = "h"
	//make(map[int][int])
	fmt.Println(m)
	var screen string
	for _, value0 := range m {
		for _, value1 := range value0 {
			screen += value1
		}
		screen += "\n"
	}
	fmt.Println(screen)
}
