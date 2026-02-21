/*
Напишите функцию объединения каналов в 1

package main

	func main() {
		chanels := make([]chan int64, 10)
		for i := range chanels {
			chanels[i] = make(chan int64)
		}

		for i := range chanels {
			go func(i int) {
				chanels[i] <- int64(i)
				close(chanels[i])
			}(i)
		}

		for v := range merge(chanels...) {
			println(v)
		}
	}
*/
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	chanels := make([]chan int64, 10)
	for i := range chanels {
		chanels[i] = make(chan int64)
	}

	for i := range chanels {
		go func(i int) {
			chanels[i] <- int64(i)
			chanels[i] <- 20
			close(chanels[i])
		}(i)
	}

	ctx, chanel := context.WithTimeout(context.Background(), time.Millisecond)
	defer chanel()

	for v := range merge(ctx, chanels...) {
		println(v)
	}
}

func merge(ctx context.Context, channels ...chan int64) chan int64 {
	res := make(chan int64)

	if len(channels) == 0 {
		close(res)
		return res
	}

	var wg sync.WaitGroup // тоже самое wg := sync.WaitGroup{}
	wg.Add(len(channels))

	for _, ch := range channels {
		go func(c chan int64) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					fmt.Println("context done")
					return
				case val, ok := <-c:
					if !ok {
						return
					}

					select {
					case res <- val:
					case <-ctx.Done():
						return
					}

				}
			}

		}(ch)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}
