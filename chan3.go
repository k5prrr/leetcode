package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 4)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		ch <- 42
	}()
	go func() {
		fmt.Println(<-ch)
		wg.Done()
		//wg.Done()
		fmt.Println(<-ch)
		wg.Done()
	}()

	wg.Wait()

}
