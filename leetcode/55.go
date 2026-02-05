/*
55. Игра в прыжки

премиум-блокировка, иконки компаний

Вам задан целочисленный массив nums.
Изначально вы располагаетесь по первому индексу массива,
и каждый элемент в массиве представляет максимальную длину вашего перехода в этой позиции.

Верните значение true, если вы можете достичь последнего индекса, или значение false в противном случае.

Пример 1:

Ввод: nums = [2,3,1,1,4]
Вывод: true
Пояснение: Переход на 1 шаг от индекса 0 к 1, затем на 3 шага к последнему индексу.

Пример 2:

Ввод: nums = [3,2,1,0,4]
Вывод: false
Пояснение: Вы всегда будете достигать индекса 3, несмотря ни на что. Максимальная длина перехода равна 0,
что делает невозможным достижение последнего индекса.

Ограничения:

	1 <= цифры.длина <= 104
	0 <= цифры[i] <= 105
*/
package main

import (
	"fmt"
)

// My Not Optim
func canJump(nums []int) bool {
	fmt.Printf("%v\n", nums)

	len0 := len(nums)
	if len0 < 2 || nums[0] >= len0 {
		return true
	}

	for i := nums[0]; i <= nums[0] && i > 0; i-- {
		if canJump(nums[i:]) {
			return true
		}
	}

	return false
}

func canJumpGpt(nums []int) bool {
	maxReach := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, i+nums[i])
		if maxReach >= n-1 {
			return true
		}
	}

	return maxReach >= n-1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	testCases := []struct {
		nums     []int
		expected bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0}, true},
		{[]int{0, 1}, false},
		{[]int{1, 1, 1, 1, 1}, true},
		{[]int{5, 0, 0, 0}, true},
		{[]int{2, 0, 0, 1}, false},
		{[]int{3, 0, 0, 0}, true},
		{[]int{1, 0, 1, 0}, false},
		{[]int{2, 0, 0}, true},
	}

	fmt.Println("Результаты проверки canJump:")
	for i, tc := range testCases {
		result := canJump(tc.nums)
		status := "OK"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("Тест %d: canJump(%v) = %v, ожидалось %v → %s\n",
			i+1, tc.nums, result, tc.expected, status)
	}
}
