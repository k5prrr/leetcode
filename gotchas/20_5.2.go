// Возврат ошибки из функции
// Верните из функции ошибку, не подключая доп пакетов

/*
 * Если метод модифицирующий, то лучше кидать обьект, а если модифицирующий, то указатель
 * но в плане хранения оно не очень.. Выбирай сам. Я за всегда указатель
 *
 * Когда возврат через интерфейсы, то всегда лучше возвращать указатель
 */

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
