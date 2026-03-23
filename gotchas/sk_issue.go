/*
wg := &sync.WaitGroup{}
wd.Add(1)
wg.Done()
wg.Wait()

гонка данных
var at atomic.Int32
at.Add(1)
at.Load()
!go run -race .

mu := &sync.Mutex{}
mur := &sync.RWMutex{} // top
mu.RLock()

	Только читатели МОМЕНТ ЧТЕНИЯ, не блочит другого читателя

mu.RUnlock()
mu.Lock()

	писатели ТОЛЬКО МОМЕНТ ЗАПИСИ!!!

mu.Unlock()

ch := make(chan int)
ch <- 1
val := <- ch

for range 100 {}

workSeconds := rand.Intn(maxWaitSeconds)
time.Sleep(time.Duration(workSeconds) * time.Second)
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var maxWaitSeconds = 4

func randWait() int {
	workSeconds := rand.Intn(maxWaitSeconds)
	time.Sleep(time.Duration(workSeconds) * time.Second)

	return workSeconds
}

func main() {
	var mainSeconds atomic.Int32
	timeStart := time.Now()

	kol := 100
	wg := &sync.WaitGroup{}
	wg.Add(kol)

	for range kol {
		go func() {
			defer wg.Done()
			mainSeconds.Add(int32(randWait()))
		}()
	}
	wg.Wait()

	fmt.Println("main:", mainSeconds.Load())
	fmt.Println("real:", time.Since(timeStart))
}

/*
 * Есть способ через каналы
 * писать в канал в рутинах
 * читать из канала столько же раз в форе
 * если не знаем сколько раз писалось, то проверка канала на закрытость
 */
