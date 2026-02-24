/*
Долгая обработка изображений
необходимо ограничить количество горутин
*/
package main

import (
	"fmt"
	"time"
	"sync"
)

type Task struct {
	ID int64
	Url string
	NewUrl string
}

func Proccess(task Task) string {
	time.Sleep(time.Second)
	return fmt.Sprintf("Файл обработан %s ID: %d\n", task.Url, task.ID)
}

func RunWorker(taskCh <- chan Task, resCh chan<- string) {
	for task := range taskCh {
		fmt.Printf("Worker start id %d\n", task.ID)
		resCh <- Proccess(task)
	}
}
func main() {
	const (
		numWorkers = 3
		numTask = 10
	)

	taskCh := make(chan Task, numTask)
	resCh := make(chan string, numTask)

	wg := sync.WaitGroup{}
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			RunWorker(taskCh, resCh)
		}()
	}

	// Генератор
	go func() {
		for i := 0; i < numTask; i++ {
			taskCh <- Task{ID:int64(i), Url:fmt.Sprintf("/%d.jpg", i)}
		}
		close(taskCh)
	}()

	go func() {
		wg.Wait()
		close(resCh)
	}()

	for res := range resCh{
		fmt.Println(res)
	}

	fmt.Println("end")
}
