package main

import "fmt"

func main() {
	// Создаем слайс с тремя элементами
	//    slice := []int{1, 2, 3}
	slice :=  make([]int, 1, 10)
	slice2 := slice
	slice = append(slice, 10)
	slice2 = append(slice2, 6)
	fmt.Println(slice)


	// Выводим длину и ёмкость слайса
	fmt.Println("Length:", len(slice))   // Длина слайса
	fmt.Println("Capacity:", cap(slice)) // Ёмкость слайса

	fmt.Println(&slice[1])
	fmt.Println(&slice2[1])

}


