package main

import (
	"log"
	"sync"
	"sync/atomic"
)

var counter int64
var counter2 atomic.Int64

func inc(wg *sync.WaitGroup) {
	defer wg.Done()

	counter++
	counter2.Add(1)
	// читает, считает, пишет
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go inc(&wg)
	}
	wg.Wait()

	log.Println(counter, counter2.Load())
}
