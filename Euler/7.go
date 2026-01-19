package main

import "fmt"

func main() {
	max := 1000
	simple := false
	final := 0
	for i := 2; i < max; i++ {
		simple = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				simple = false
				break
			}
		}
		if !simple {
			continue
		}
		final = i
		//fmt.Println(i)
	}
	fmt.Println(final)
}
