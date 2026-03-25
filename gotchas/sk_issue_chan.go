/*
https://www.youtube.com/watch?v=k-1OEYl7N8Q

Каналы как 2 комнаты
в одном producers
в другом consumers

А буфер как количество окошек..
И кто-то (с любой стороны) постоянно занимает окошки, остальные ждут..

Канал всегда инициализируем!
ch := make(chan int)

go func(){
	wg.Wait()
	close(ch)
}()
Но канал НЕ всегда надо закрывать





chan НЕбуф 	| Открытый			| Закрытый		| Неинициализированный
Чтение		| L Блок до писателя| V Zero val	| ! Блок Навсегда
Запись		| L Блок до читателя| ! panic		| ! Блок Навсегда
Закрытие	| V Закроется		| ! panic		| ! panic


chan буф 	| Открытый частичн полн	| Открытый полн		| Открытый пустой	| Закрытый частичн полный
Чтение		| V	чтение значения		| V	чтение значения	| L Блок до писателя| V	чтение значения
Запись		| V запись значения		| L Блок до читателя| V запись значения	| ! panic


<-chan Только на чтение
Запись		| ! Ошибка компиляции
Закрытие	| ! Ошибка компиляции

chan<- Только на запись
Чтение		| ! Ошибка компиляции


планировщик запускает ту горутину, которая в очереди последняя (стек)
*/

package main

import (
	"fmt"
	"time"
)

// Паттерн Генератор
func writer() <-chan int {
	ch := make(chan int)

	go func() {
		for i := range 9 {
			ch <- i + 1
			time.Sleep(500 * time.Millisecond)
		}
		close(ch)
	}()

	return ch
}

func dobbler(ch <-chan int) <-chan int {
	res := make(chan int)
	go func() {
		for v := range ch {
			res <- v * 2
		}
		close(res)
	}()

	return res
}

func reader(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
		//time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	ch := writer()

	for {
		v, ok := <-ch
		if !ok {
			break
		}

		fmt.Println("v =", v)
	}
	/*
		OR
		for v:= range ch {
			fmt.Println("v =", v)
		}

		OR
		for range 10 {
			v:= <- ch
			fmt.Println("v =", v)
		}
	*/

	// Паттерн пайплайн или билдер
	reader(dobbler(writer()))
}
