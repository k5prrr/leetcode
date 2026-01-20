/*

 Если больше чем копасити, то ошибка
 Если как окпасити, то дополнение пустыми значениями
 data3[1:3] Тут копасити не пересоздаётся, а от него отнимается столько же сколько с начала (только когда в начале вырезаешь)
при вырезании с конца копасити остаётся тот же

ещё мы не можем начать вырезать с конца, тип если копасити 10, то вырез с 10 вызывает ошибку

data4 := data3[9:9]
data4 [], len 0, cap 1

минимально капасити 1, его не может быть 0
а, не, может, если так же вырезать
*/package main

import (
	"fmt"
)


func sliceCapacityDemo(data []int, start int) []int {
	subSlice := data[start:] // Тут копасити не пересоздаётся, а от него отнимается столько же (только когда в начале вырезаешь)
	return subSlice
}

func main() {
	data := make([]int, 5, 10)
	for i := range data {
		data[i] = i + 1
	}

	fmt.Printf("Изначально %v, len %d, cap %d\n", data, len(data), cap(data))

	data = sliceCapacityDemo(data, 1)
	fmt.Printf("После sliceCapacityDemo %v, len %d, cap %d\n", data, len(data), cap(data)) // 2 3 4 5, leb4, cap9


	data2 := make([]int, 0, 3)
	data2 = sliceCapacityDemo(data, 2)
	fmt.Printf("После sliceCapacityDemo %v, len %d, cap %d\n", data2, len(data2), cap(data2)) // 4 5, leb2, cap7


// Свои эксперименты
	data3 := make([]int, 0, 0)
	data4 := data3[:]
	fmt.Printf("data4 %v, len %d, cap %d\n", data4, len(data4), cap(data4)) // 4 5, leb2, cap7

}
