/* Интерфейсы

приведение типов - дорого
mane, ok := dfdfdf.(int)


надо всё проверять на ошибки!
*/

package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func (c *Cache) Store(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}
func (c *Cache) Load(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.data[key]
	return val, ok
}

func main() {
	cache := &Cache{
		data: make(map[string]interface{}),
	}

	cache.Store("name", "Alice")
	cache.Store("age", 25)

	// 1
	nameI, ok := cache.Load("name")
	if !ok {

	}
	name, ok := nameI.(string)
	if !ok {

	}
	fmt.Println("Name: ", name)

	// 2
	ageI, _ := cache.Load("age")
	age, _ := ageI.(int)
	fmt.Println("Age: ", age)

	// 3
	height := cache.Load("height")
	if height == nil {
		fmt.Println("height not found")
	}

	// 4
	width := cache.Load("width").(float64)
	fmt.Println("Width: ", width)

}
