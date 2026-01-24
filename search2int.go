/*
Поиск двух чисел с заданной суммой
Программа находит два числа в массиве, сумма которых равна заданному значению. Используется хэш-таблица для оптимизации поиска.

 */
package main

import "fmt"

func count1(a []int, sum int) [2]int {
	m := make(map[int]int)
	for index, value := range a {
		m[value] = index
	}

	for index, value := range a {
		mindex, exists := m[sum-value]
		if exists && index != mindex {
			return [2]int{index, mindex}
		}
	}

	return [2]int{0, 0}
}

func main() {
	fmt.Println(count1([]int{3, 8, 2, 2, 4}, 6))
}
