/*
Дан массив из трёх целых чисел и целевое значение sum.
Найдите два различных индекса в массиве таких, что сумма элементов по этим индексам равна заданному значению sum.
Если такие индексы существуют, верните их в виде массива из двух элементов.
Если подходящей пары не существует, верните массив [-1, -1].

Гарантируется, что массив содержит ровно три элемента.

Пример:
Вход: a = [3, 2, 4], sum = 6
Выход: [1, 2] (поскольку a[1] + a[2] = 2 + 4 = 6)
*/
package main

import (
	"fmt"
)

//var a = [3]int{3,2,4}

func main() {
	//sum := 6

	fmt.Println(count1([3]int{3,2,4}, 6));

}

func count1(a [3]int, sum int) [2]int {
	fmt.Println(a)
	m := make(map[int]int)
	for index, value := range a {
		m[value] = index
	}
	fmt.Println(m)

	for index, value := range a {
		mindex , exists := m[sum-value]
		if (exists && index != mindex) {
			return [2]int{index, mindex}
		}
	}

	return [2]int{-1,-1}
}
