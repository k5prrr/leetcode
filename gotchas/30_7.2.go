/*
	напишите реализацию InMemory кэша

package main

	type Cache interface {
		Set(key, value string)
		Get(key string) (string, bool)
	}
*/
package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Cache interface {
	Set(key, value string)
	Get(key string) (string, bool)
}

type Shard struct {
	data map[string]string
	mu   sync.RWMutex
}
type MyCache struct {
	shards []*Shard
}

func NewMyCache(initSize int64) *MyCache {
	shards := make([]*Shard, initSize)

	for i := range shards {
		shards[i] = &Shard{
			data: make(map[string]string),
		}
	}

	return &MyCache{
		shards: shards,
	}
}

func (c *MyCache) getShard(key string) *Shard {
	hasher := fnv.New32a()
	_, _ = hasher.Write([]byte(key))
	hash := hasher.Sum32()

	return c.shards[hash%uint32(len(c.shards))]
}

func (c *MyCache) Set(key, value string) {
	shard := c.getShard(key)

	shard.mu.Lock()
	defer shard.mu.Unlock()

	shard.data[key] = value
}

func (c *MyCache) Get(key string) (string, bool) {
	shard := c.getShard(key)

	shard.mu.RLock()
	defer shard.mu.RUnlock()

	val, ok := shard.data[key]

	return val, ok
}

func main() {
	cache := NewMyCache(1000)

	const (
		writerCount = 50  // рутин на запись
		readerCount = 200 // рутин на чтение
		duration    = 5 * time.Second
		maxKeys     = 1000
	)

	var wg sync.WaitGroup
	var writeOps, readOps, errors int64

	done := make(chan struct{})

	// записи
	for i := 0; i < writerCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				default:
					key := fmt.Sprintf("key_%d", rand.Intn(maxKeys))
					value := fmt.Sprintf("value_%d_%d", id, rand.Intn(10000))

					cache.Set(key, value)
					atomic.AddInt64(&writeOps, 1)

					time.Sleep(time.Microsecond * time.Duration(rand.Intn(50)))
				}
			}
		}(i)
	}

	// чтение
	for i := 0; i < readerCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				default:
					key := fmt.Sprintf("key_%d", rand.Intn(maxKeys))

					val, ok := cache.Get(key)
					atomic.AddInt64(&readOps, 1)

					// проверка целостности
					if ok && len(val) == 0 {
						atomic.AddInt64(&errors, 1)
						fmt.Printf("ERROR: Пустое значение для существующего ключа %s\n", key)
					}

					time.Sleep(time.Microsecond * time.Duration(rand.Intn(10)))
				}
			}
		}(i)
	}

	go func() {
		time.Sleep(duration)
		close(done)
	}()
	wg.Wait()

	fmt.Printf("\nДлительность:        %v\n"+
		"Горутины (запись):   %d\n"+
		"Горутины (чтение):   %d\n"+
		"Операций записи:     %d\n"+
		"Операций чтения:     %d\n"+
		"Ошибок целостности:  %d\n",
		duration,
		writerCount,
		readerCount,
		atomic.LoadInt64(&writeOps),
		atomic.LoadInt64(&readOps),
		atomic.LoadInt64(&errors),
	)

	if atomic.LoadInt64(&errors) == 0 {
		fmt.Println("Тест пройден успешно!")
	} else {
		fmt.Println("Обнаружены ошибки!")
	}
}
