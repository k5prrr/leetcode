/*
 * Чтение из одного и того же канала в разных горутинах
 */
package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(10_000 * 2)

	// Писари
	go func() {
		for i := range 10_000 {
			defer wg.Done()
			ch <- i
		}
	}()
	go func() {
		for i := range 10_000 {
			defer wg.Done()
			ch <- i
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	// Читатели
	go func() {
		for v := range ch {
			fmt.Println("v=", v, " worker 1")
		}
	}()

	go func() {
		for v := range ch {
			fmt.Println("v=", v, " worker 2")
		}
	}()
	for v := range ch {
		fmt.Println("v=", v, " worker 3")
	}

}
