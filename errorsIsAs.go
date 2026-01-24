/*
 * Пример 1: ErrNotFound пойман через Is
 * Пример 2: ErrTimeout пойман через Is внутри wrapped ошибки
 * Пример 3: найдена кастомная ошибка с сообщением: что-то сломалось
 * Пример 4: Is не нашёл ErrTimeout — как и ожидалось
 *
 */
package main

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")
	ErrTimeout  = errors.New("timeout")
)

type MyError struct {
	msg string
}
func (e *MyError) Error() string {
	return "my error: " + e.msg
}

func main() {
	err1 := ErrNotFound
	if errors.Is(err1, ErrNotFound) {
		fmt.Println("Пример 1: ErrNotFound пойман через Is")
	}

	// ошибка с оборачиванием
	err2 := fmt.Errorf("произошла ошибка: %w", ErrTimeout)
	if errors.Is(err2, ErrTimeout) {
		fmt.Println("Пример 2: ErrTimeout пойман через Is внутри wrapped ошибки")
	}

	// Пример 3: кастомная ошибка и As
	err3 := &MyError{msg: "что-то сломалось"}  // &MyError{"что-то сломалось"} тоже самое
	wrappedErr := fmt.Errorf("внутри функции произошло: %w", err3)

	var target *MyError
	if errors.As(wrappedErr, &target) {
		fmt.Printf("Пример 3: найдена кастомная ошибка с сообщением: %s\n", target.msg)
	}

	// Пример 4: Is не сработает на другой тип
	if errors.Is(wrappedErr, ErrTimeout) {
		fmt.Println("Пример 4: Неожиданно — это не выполнится")
	} else {
		fmt.Println("Пример 4: Is не нашёл ErrTimeout — как и ожидалось")
	}
}
