/*
 a1 := []int{1, 2, 3, 4, 5}
 go run Koz10_3.2.go
 [1 2 3 4 5] [1 2 3 4 5 6] [1 2 3 4 5 7]
cap растёт в 2 раза до 256


[admin@arch leetcode]$ go run Koz10_3.2.go
[1 2 3 4 5] [1 2 3 4 5 7 8] [1 2 3 4 5 7]

 */
package main

import (
	"fmt"
)

func main() {
	//a1 := []int{1, 2, 3, 4, 5}
	//OR
	a1 := make([]int, 0, 10)
	a1 = append(a1, []int{1, 2, 3, 4, 5}...)

	a2 := append(a1, 6, 8)
	a3 := append(a1, 7)
	fmt.Println(a1, a2, a3)
}
