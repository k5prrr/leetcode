package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	close(ch3)

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- 1
	}()

	timer := time.NewTimer(3 * time.Second)

	// Всегда лучше использовать контекст
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select { // Он блокирующий, если нет дефолт
	case v := <-ch1:
		fmt.Println("v=", v, " ch1")
	case v := <-ch2:
		fmt.Println("v=", v, " ch2")
	case v := <-ch3: // Чтение закрытого вернёт zer
		fmt.Println("v=", v, " ch3")

	case t := <-time.After(4 * time.Second):
		fmt.Println("v=", t, " time")

	case t := <-timer.C:
		fmt.Println("v=", t, " timer ")

	case s := <-ctx.Done():
		fmt.Println("v=", s, " timer ")

	}

}
