/*
 https://euler.jakumo.org/problems.html
 Нахождение наибольшего простого делителя
 Программа находит наибольший простой делитель числа 600851475143
 */
package main

import "fmt"

func main() {
	n := 600851475143
	i := 1

	for n != 1 {
		i++
		if n%i != 0 {
			continue
		}
		n /= i
	}
	fmt.Println(i)
}
