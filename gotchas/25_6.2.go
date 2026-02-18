package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

/*
	func getCount() float64 {
		time.Sleep(3 * time.Second)
		return 5.3
	}

	func main() {
		count := getCount()
		fmt.Printf("Скидка %v\n", count)
	}
*/
var defaultTimeout = time.Second

func getCount() float64 {
	time.Sleep(3 * time.Second)
	return 5.3
}
func modGetCount(ctx context.Context) (float64, error) {
	// Проверяем исправляем контекст
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
		defer cancel()
	}

	// Кидаем туда данные
	ch := make(chan float64)
	go func() {
		ch <- getCount()
		close(ch)
	}()

	// Кто быстрее
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case res := <-ch:
		return res, nil
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := modGetCount(ctx)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Скидка %v\n", count)
}
