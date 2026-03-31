package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 80*time.Microsecond)
	defer cancel()

	ch := make(chan int)

	go func() {
		for i := range 1000 {
			select {
			case ch <- i:
			case <-ctx.Done():
				break
			}

		}
		close(ch)
	}()

	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(v)
		case <-ctx.Done():
			return
		}
	}

}
