// Ещё не доделал..
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processData(val int) int {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

	return val * 2
}

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for i := range 10 {
			in <- i
		}
		close(in)
	}()

	now := time.Now()
	processParallel(in, out, 5)

	for val := range out {
		fmt.Println(val)
	}
	fmt.Println(time.Since(now))
}

// + операция должна выполняться не более 5 секунд
func processParallel(in <-chan int, out chan<- int, numWorkers int) {
	sem := make(chan struct{}, numWorkers)

	for v := range in {
		sem <- struct{}{}

		go func() {
			out <- processData(v)
		}()

		<-sem
	}
	close(sem)
	close(out)
}
