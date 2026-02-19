/*
Ревью кода

package main

import (

	"fmt"
	"math/rand"
	"sync"
	"time"

)

	func LongCalc(n int) int {
		seckondToSleep := rand.Float64() * float64(n)
		time.Sleep(time.Duration(seckondToSleep))
		return n + 1
	}

var cache = map[int]int{}

	func CachedLongCalc(n int) int {
		var mu sync.Mutex

		mu.Lock()
		found, ok := cache[n]
		mu.Unlock()

		if !ok {
			val := LongCalc(n)

			mu.Lock()
			cache[n] = val
			mu.Unlock()

			return val
		}

		mu.Unlock()

		return found
	}

	func main() {
		nums := []int{5, 10, 22, 5, 5, 22, 10, 3}
		for _, n := range nums {
			val := CachedLongCalc(n)
			fmt.Printf("%d => %d\n", n, val)
		}
	}



	need len Lock == len Unlock !
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cache struct {
	mu   sync.RWMutex
	data map[int]int // При переполнении памяти Game over
} // Решать очередью или по таймеру(с контролем таймаутов) или по эфективности(хит мис атомики)
func NewCache() *Cache {
	return &Cache{
		data: make(map[int]int),
	}
}

var cache = NewCache()

func CachedLongCalc(n int) int {
	cache.mu.RLock()
	found, ok := cache.data[n]
	cache.mu.RUnlock()

	if !ok {
		val := LongCalc(n)

		cache.mu.Lock()
		// Тут можно проверить существование ключа второй раз
		cache.data[n] = val
		cache.mu.Unlock()

		return val
	}

	return found
}

func LongCalc(n int) int {
	seckondToSleep := rand.Float64() * float64(n)
	time.Sleep(time.Duration(seckondToSleep))
	return n + 1
}

func main() {
	nums := []int{5, 10, 22, 5, 5, 22, 10, 3}

	wg := sync.WaitGroup{}
	wg.Add(len(nums))

	for _, n := range nums {
		go func() {
			defer wg.Done()

			val := CachedLongCalc(n)
			fmt.Printf("%d => %d\n", n, val)
		}()
	}

	wg.Wait()
}
