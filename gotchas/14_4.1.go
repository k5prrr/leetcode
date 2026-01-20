/*
	мапы
	это ассоциативный массив
	ключ + значение

мапа периодически выделяет память и делает релокацию, как слайсы

В мапе данные хранятся в виде бакетов
бакет - структура данных, которая хранит в себе по 8 пар ключ значение

ключ пропускается через хэш, где на выходе большое число
это число делится на количество бакетов и остаток является номером бакета

	в каждом бакете по 8 элементов

если переполнение, то происходит как связанный список

когда в среднем по бакету 6.5 занятой ёмкости, то релокация/эвакуация
поэтапно

Мьютексы только 2 типов:
1 обычные и
2 rw/ чтение запись

обычные
lock() - захватывает мьютекс (если занят — поток блокируется).
unlock() - освобождает мьютекс.
try_lock() - пытается захватить мьютекс без блокировки; возвращает true, если успешно, иначе false

rw
несколько потоков могут читать одновременно
только один поток может писать, и при этом никто не может читать или писать
lock_shared() / rdlock() - захватить на чтение.
unlock_shared() / rdunlock() - освободить чтение.
lock() / wrlock() - захватить на запись.
unlock() / wrunlock() - освободить запись.
try_lock_shared(), try_lock()

инициализация мапы в обязаловку
*/
package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	data map[string]string
	m    sync.RWMutex
}

// Конструктор (использовать всегда для создания объектов)
func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		data: make(map[string]string),
	}
}

func (c *ConcurrentMap) GetOrCreate(key, value string) string {
	c.m.RLock()
	val, ok := c.data[key]
	c.m.RUnlock()

	if ok {
		return val
	}

	c.m.Lock()
	defer c.m.Unlock()

	if val, ok := c.data[key]; ok {
		return val
	}
	c.data[key] = value

	return value
}

func main() {
	cm := NewConcurrentMap()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		val := cm.GetOrCreate("k1", "v1")
		fmt.Println("routin1 ", val)
	}()

	go func() {
		defer wg.Done()

		val := cm.GetOrCreate("k1", "v2")
		fmt.Println("routin2 ", val)
	}()

	wg.Wait()
}
