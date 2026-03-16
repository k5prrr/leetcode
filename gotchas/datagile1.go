package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func longFunc() int64 {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	return 11
}

// ---

func modiLongFunc(ctx context.Context) (int64, error) {
	ch := make(chan int64, 1) // Буфиризированный, чтоб избежать утечку горутин

	go func() {
		res0 := longFunc()
		select {
		case <-ctx.Done():
			return
		case ch <- res0:
			close(ch)
		}
	}()

	select {
	case <-ctx.Done():
		return 0, fmt.Errorf("ctx end: %s", ctx.Err())
	case res := <-ch:
		return res, nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	res, err := modiLongFunc(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Good res: %d\n", res)
}

// ---

// Best
func longFuncTimer(ctx context.Context) (int64, error) {
	timer := time.NewTimer(time.Duration(rand.Intn(5)) * time.Second)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case <-timer.C:
		return 11, nil
	}
}
