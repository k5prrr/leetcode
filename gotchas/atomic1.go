package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var counter int64

	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
		}()
	}

	fmt.Scanln()
	fmt.Println("Counter:", counter)
}
