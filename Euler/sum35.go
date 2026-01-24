/*
 Сумма чисел, кратных 3 или 5
 Программа вычисляет сумму всех чисел меньше 1000, которые кратны 3 или 5. Это классическая задача из Project Euler.
 */
package main

import "fmt"

func main() {
	sum := 0

	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
