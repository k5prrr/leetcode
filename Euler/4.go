/*
 https://euler.jakumo.org/problems.html
 Палиндромы
 Программа находит наибольший палиндром, который является произведением двух трёхзначных чисел.
 */
package main

import "fmt"

func isPalindrome(n int) bool {
	s := fmt.Sprintf("%d", n)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	maxPalindrome := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			product := i * j
			if isPalindrome(product) && product > maxPalindrome {
				maxPalindrome = product
			}
		}
	}
	fmt.Println(maxPalindrome)
}
