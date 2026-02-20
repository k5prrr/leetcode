/* Найти и исправить проблемы в коде ниже
У парковки ограниченное количество мест
программа должна отработать корректно и завершиться без зависаний

package main

import (
	"fmt"
	"sync"
	"time"
)

type ParkingLot struct {
	slots int64
}

func (p *ParkingLot) Park(carID int64) {
	fmt.Printf("Машинка %d паркуется\n", carID)
	time.Sleep(time.Second)
	fmt.Printf("Машинка %d припаркована\n", carID)
}

func main() {
	parking := &ParkingLot{slots: 3}

	carIDs := []int64{1, 2, 3, 4, 5, 6, 7}
	var wg sync.WaitGroup
	wg.Add(len(carIDs))

	for _, carID := range carIDs {
		go func(id int64) {
			defer wg.Done()

			parking.Park(id)
		}(carID)
	}

	wg.Wait()
	fmt.Println("Всё")

}
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type ParkingLot struct {
	slots chan struct{}
	count atomic.Int64
}

func NewParkingLot(slots int64) *ParkingLot {
	newPark := &ParkingLot{
		slots: make(chan struct{}, slots),
	}
	newPark.count.Store(slots)

	return newPark
}

func (p *ParkingLot) Park(carID int64) {
	p.slots <- struct{}{}
	// OR
	// if p.count.Load() == 0 { нет мест return }
	// p.count.Add(-1)

	fmt.Printf("Машинка %d паркуется\n", carID)
	time.Sleep(time.Second)
	fmt.Printf("Машинка %d уехала\n", carID)

	// p.count.Add(1)
	<-p.slots
}

func main() {
	parking := NewParkingLot(3)

	carIDs := []int64{1, 2, 3, 4, 5, 6, 7}
	var wg sync.WaitGroup
	wg.Add(len(carIDs))

	for _, carID := range carIDs {
		go func(id int64) {
			defer wg.Done()

			parking.Park(id)
		}(carID)
	}

	wg.Wait()
	fmt.Println("Всё")
}
