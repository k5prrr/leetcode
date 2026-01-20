/*
пустая структура ничего не весит



*/
package main

import (
	"fmt"
	//"crypto/rand" сильнее
	"math/rand"
)


func main() {
	fmt.Println(uniq2(10))
}

func uniq(n int) []int {
	result := make([]int, 0, n)
	for i:= 0 ; i<n; i++ {
		result = append(result, i)
	}
	return result
}

func uniq2(n int) []int {
	m := make(map[int]struct{}, n)
	res := make([]int, 0, n)

	for len(res) < n {
		//tmp := rand.Int()
		tmp := rand.Intn(128)

		if _, ok := m[tmp]; !ok {
			res = append(res, tmp)
			m[tmp] = struct{}{} // struct{} - формат, {} - как содержимое объекта
		}
	}

	return res
}
