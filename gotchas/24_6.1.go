/*
 * Доработать SimulateRequest
 * - чтоб код работал в конкурентной среде
 * - при долгом ожидании код отваливался по таймауту
 * - печаталось время выполнения запроса
var counter int64
func SimulateRequest() int64 {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	counter++

	return counter
}
 func main() {
	val := SimulateRequest()
	log.Printf("Счётчик %d\n", val)
 }
*/

package main

import (
	"context"
	"log"
	"math/rand"
	"sync/atomic"
	"time"
)

var counter atomic.Int64

func SimulateRequest(ctx context.Context) (int64, error) {
	// Таймер для функции
	timeStart := time.Now()
	defer func() {
		log.Printf("Время выполнения запроса: %v\n", time.Since(timeStart))
	}()

	ch := make(chan int64)
	go func() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		counter.Add(1)
		ch <- counter.Load()
		close(ch)
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case count := <-ch:
		return count, nil
	}
}
func main() {
	log.Println("Start")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	val, err := SimulateRequest(ctx)
	if err != nil {
		log.Printf("Ошибка выполнения запроса %v\n", err)
	}

	log.Printf("Счётчик %d\n", val)
}
