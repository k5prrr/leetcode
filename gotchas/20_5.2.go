// Возврат ошибки из функции
// Верните из функции ошибку, не подключая доп пакетов

package main

func main() {
	println(handle())
}

type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}

func handle() error {
	return &CustomError{message: "Ошибка"}
}
