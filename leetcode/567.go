package main

import (
	"fmt"
	//"strings"
)

func main() {
	str := "aab"
	text := "aaarabrabaccd"

	strlen := len(str)
	result := 0
	//arr0 := strings.Split(str, "")

	map0 := make(map[rune]int)
	for _, value := range str {
		_, exist := map0[value]
		if exist {
			map0[value]++
			continue
		}
		map0[value] = 1
	}
	fmt.Println(map0)
	fmt.Println(strlen)

	for index, value := range text {
		fmt.Println("Pluse", index, text[index])

		_, exist := map0[value]
		if exist {
			map0[value]--
			if map0[value] >= 0 {
				result++
			} else {
				result--
			}
		}

		if index >= strlen {
			valueOld := rune(text[index-strlen])
			fmt.Println("Minus", index-strlen, valueOld)
			_, existOld := map0[valueOld]
			if existOld {
				map0[valueOld]++
				if map0[valueOld] > 0 {
					result--
				} else {
					result++
				}
			}
		}

		if result == strlen {
			fmt.Println("final", true)
			return
		}
	}
	fmt.Println("final", false)
}
