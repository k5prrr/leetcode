/*
 * Часто на собесах
 * Написать обёртку для функции которая долго работает
 */
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func randomTimeWork() int {
	res := rand.Intn(5)
	time.Sleep(time.Duration(res) * time.Second)

	return res
}

// ---
func wrapperRandomTimeWork(ctx context.Context) (int, error) {
	ch := make(chan int)
	go func() {
		select {
		case ch <- randomTimeWork():
		case <-ctx.Done():
		}
		close(ch)
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := wrapperRandomTimeWork(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("good: ", res)
}
