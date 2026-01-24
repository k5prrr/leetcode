package main
// Lock-Free очередь на атомиках

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type Queue struct {
	head uint64 // индекс для записи
	tail uint64 // индекс для чтения
	data []interface{}
}

func NewQueue(size int) *Queue {
	return &Queue{
		data: make([]interface{}, size),
	}
}

// Push добавляет элемент в очередь
func (q *Queue) Push(value interface{}) bool {
	for {
		head := atomic.LoadUint64(&q.head)
		tail := atomic.LoadUint64(&q.tail)
		if head - tail >= uint64(len(q.data)) {
			// Очередь заполнена
			return false
		}
		if atomic.CompareAndSwapUint64(&q.head, head, head+1) {
			q.data[head % uint64(len(q.data))] = value
			return true
		}
	}
}

// Pop извлекает элемент из очереди
func (q *Queue) Pop() (interface{}, bool) {
	for {
		tail := atomic.LoadUint64(&q.tail)
		head := atomic.LoadUint64(&q.head)
		if tail >= head {
			// Очередь пуста
			return nil, false
		}
		val := q.data[tail % uint64(len(q.data))]
		if atomic.CompareAndSwapUint64(&q.tail, tail, tail+1) {
			return val, true
		}
	}
}

func main() {
	q := NewQueue(4)

	q.Push("A")
	q.Push("B")
	q.Push("C")

	fmt.Println(q.Pop()) // A
	fmt.Println(q.Pop()) // B
	fmt.Println(q.Pop()) // C
}
