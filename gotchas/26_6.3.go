/*
	func main() {
		ch := make(chan bool)
		go func() {
			time.Sleep(3 * time.Second)
			fmt.Println("Проснулась корутинка")
			ch <- false
		}()

		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				fmt.Println("Тик")
				ch <- true // Тут дедлок так ка никто не прочитает, для исправления добавить буфер
			case value := <-ch:
				fmt.Printf("Получено: %t\n", value)
				return
			}
		}
	 }
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan bool, 1)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Проснулась корутинка")
		ch <- false
		close(ch)

		wg.Done()
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Тик")
			//ch <- true
		case value := <-ch:
			fmt.Printf("Получено: %t\n", value)
			wg.Wait()
			return
		}
	}
}

/*
 * Проверить закрыт ли канал можно только прочитав
 *
 * НО Если канал закроется между проверкой и записью, программа упадет.
 *
 * Канал должен закрывать только тот, кто в него пишет
 *
select {
case val, ok := <-ch:
    if !ok {
        // Канал закрыт и пуст
        fmt.Println("Канал закрыт")
    } else {
        // Получили значение
        fmt.Println("Получено:", val)
    }
default:
    // Канал открыт, но сейчас пуст (чтение заблокировалось бы)
    // ИЛИ в канале есть данные, но мы их не взяли (редкий кейс race condition)
    fmt.Println("Канал открыт, но данных нет прямо сейчас (или не блокируемся)")
}



Вот способ лучше
Как избежать паники? Следить за состоянием через флаги или context

func worker(ctx context.Context, ch chan<- int) {
    for {
        select {
        case <-ctx.Done():
            // Сигнал на остановку
            close(ch) // Закрываем канал, так как мы отправитель
            return
        case ch <- generateData():
            // Пишем данные
        }
    }
}


*/
