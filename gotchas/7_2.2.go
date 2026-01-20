/*
ascii символы
2^8 в степени 8 = 256

1 байт - 8 бит = 256 вариантов
2 байта - 16 бит = 256*256 = 65_536 вариантов
3 256*256*256
4 = 4_294_967_296 вариантов (Четыре миллиарда, двести девяносто четыре миллиона ..)

Руна - от 1 до 4 байт (зависит от того где в таблице находится символ, чем ближе к 0, тем меньше)

rune - тоже самое, что int32

len возвращает количество байт

Алиас - имя которое используется вместо основного


fmt.Printf, %c
это глагол форматирования  (format verb),
используемый для вывода символа , соответствующего указанному коду Unicode (rune) .

*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "ddЯЙ漢"                                                            // 1+1+ 2+2+ 3 = 9
	fmt.Println("Длина через len:", len(str))                                  // итог 9
	fmt.Println("Длина через RuneCountInString:", utf8.RuneCountInString(str)) // 5

	for i := 0; i < len(str); i++ {
		fmt.Printf("%d %d %c \n", i, str[i], str[i])
	}
	fmt.Println("Первые")
	for i, r := range str {
		fmt.Printf("%d %d %s \n", i, r, string(str[i]))
	}
	fmt.Println("Полные")
	for i, r := range []rune(str) {
		fmt.Printf("%d %d %s \n", i, r, string(r))
	}
}
