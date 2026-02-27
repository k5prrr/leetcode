/*
симафор
Делаете систему для загрузки файлов с удалённого сервера
каждый файл грузится через отдельное соединение
сервер ограничивает количество одновременных соединений
в программе запрещено запускать фиксированное количество горутин
нужно сделать динамически изменяемое ограничение
количества одновременно работающих соединений
*/

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func downloadFile(url string) {
	fmt.Printf("Url %s\n", url)
	randomSleep := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(randomSleep)
	fmt.Printf("Downloaded url %s\n", url)
}

func main() {
	gorutinesLimit := 3

	files := []string{"file1", "file2", "file3", "file4", "file5", "file6", "file7"}

	wg := sync.WaitGroup{}
	wg.Add(len(files))

	semaphore := make(chan struct{}, gorutinesLimit)

	for _, url := range files {
		semaphore <- struct{}{}
		go func() {
			defer func() {
				<-semaphore
				wg.Done()
			}()

			downloadFile(url)
		}()
	}

	fmt.Println(runtime.NumGoroutine()) // 4 Так как - основной поток main
	wg.Wait()
	fmt.Println("end")
}
