package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, "=")
	var result int64 = 0
	for _, num := range nums {
		result += int64(num)
	}
	fmt.Println(result)
}

func main() {
	sum()
	sum(1)
	sum(1, 2)
}
