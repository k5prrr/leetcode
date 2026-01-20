/*
 Слайсы
 структура



 [admin@arch leetcode]$ go run Koz9_3.1.go
 Слайс старт [10 20 30 40]
 Слайс в функции [10 20 50 60]
 Слайс после изменений [10 20 50 60]

 А если бы добавился бы 70, то была бы релокация и первичный слайс не поменялся
 так как добавление сразу
 [admin@arch leetcode]$ go run Koz9_3.1.go
 Слайс старт [10 20 30 40]
 Слайс в функции [10 20 50 60 70]
 Слайс после изменений [10 20 30 40]
 [admin@arch leetcode]$

 [admin@arch leetcode]$ go run Koz9_3.1.go
 Слайс старт [10 20 30 40]
 Слайс в функции [10 20 50 60 70]
 Слайс после изменений [10 20 50 40]


 */

package main

import (
	"fmt"
)

func main() {
	data := []int{10, 20, 30, 40}
	fmt.Println("Слайс старт", data) // 10, 20, 30, 40

	modify(data[:2])
	fmt.Println("Слайс после изменений", data) // 10, 20, 50, 40
}

func modify(slice []int) {
	//slice = append(slice, 50, 60)
	//slice = append(slice, 50, 60, 70)
	slice = append(slice, 50)
	slice = append(slice, 60, 70)
	fmt.Println("Слайс в функции", slice) // 10, 20, 50, 60
}
