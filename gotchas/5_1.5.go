/*
Массив с количеством это как отдельный тип данных

слайс - структура с указателями на область памяти
*/
package main

import (
	"fmt"
)

func modifyArray(arr [3]int) {
	arr[0] = 10
	fmt.Println("modifyArray result: ", arr) // 10 2 3
}
func modifySlice(slice []int) {
	slice[0] = 10
	fmt.Println("modifySlice result: ", slice)
}

func main() {
	array := [3]int{1, 2, 3} // чтоб не менять цифру, можно [...]int{1, 2, 3}
	slice := array[:]        // Та же самая область памяти

	//	fmt.Println(reflect.TypeOf(array))
	//fmt.Println(array[1:]) // 2 3 С какого элемента мы берём
	//fmt.Println(array[1:1])  Пустой массив
	//	fmt.Println(array[1:3:8])  8 - ёмкость слайса-1 и она не должна превышать массив
	// [1 с какого включая : 3 по какое не включая :8 копасити-сКакого = 7 ] так как по умолчанию копасити вытаскивается почти с размером массива
	// Всегда ставь  во 2 и 3 одинак значен

	fmt.Println("Before array: ", array) // 1 2 3
	modifyArray(array)                   // 10 2 3
	fmt.Println("After array: ", array)  // 1 2 3

	fmt.Println("Before slice: ", slice) // 1 2 3
	modifySlice(slice)                   // 10 2 3
	fmt.Println("After slice: ", slice)  // 10 2 3

	fmt.Println("Final array: ", array) // 10 2 3

	array2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := array2[2:4:4]
	fmt.Println("cap ", cap(slice2)) // cap  2
}

/*
Before array:  [1 2 3]
modifyArray result:  [10 2 3]
After array:  [1 2 3]

Before slice:  [1 2 3]
modifySlice result:  [10 2 3]
After slice:  [10 2 3]

Final array:  [10 2 3]
*/
