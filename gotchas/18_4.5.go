/*
 * Про синк мапу
 *
 * синк мапа под капотом
 *
 type Map struct {
 	mu Mutex
  	read atomic.Value // read only (кэш)
   dirty map[interface{}]*entry
   misses int
 }




*/
// Представьте, что у вас есть система, в которой нужно обрабатывать множество запросов,
// и для каждого из них требуется выполнить дорогостоящую операцию
// (например, запрос к базе данных или сложные вычисления).
// Чтобы не выполнять повторно одинаковую операцию для одного и того же ключа,
// вы хотите использовать кэш.

// Написать функцию GetOrCompute, которая принимает ключ,
// функцию для вычисления значения, и возвращает либо значение из кэша, либо результат вычислений

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache sync.Map

func GetOrCompute(key string, computeFunc func() string) string {
	if val, ok := cache.Load(key); ok {
		fmt.Printf("Данные %s из кэша\n", key)
		return val.(string)
	}

	newVal := computeFunc()

	cache.Store(key, newVal)

	return newVal
}

func calculation(userID string) string {
	time.Sleep(time.Second)
	fmt.Printf("Вычислили %s\n", userID)

	return fmt.Sprintf("[%s]", userID)
}

func main() {
	userIDs := []string{"user1", "user2", "user3", "user1", "user2", "user2", "user2", "user2", "user3"}

	wg := sync.WaitGroup{}

	for _, userID := range userIDs {
		wg.Add(1)

		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

			res := GetOrCompute(userID, func() string {
				return calculation(userID)
			})

			fmt.Println(res)
		}()
	}
	wg.Wait()
}
