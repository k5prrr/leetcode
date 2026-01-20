package main

import "fmt"

func main() {
	f := func(i int) func() int {
		return func() int {
			i++
			return i
		}
	}
	g := f(0)
	fmt.Println(g(), g())
}
