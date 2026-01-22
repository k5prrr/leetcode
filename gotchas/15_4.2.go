/*
 * Мапы 2
 * */

package main

import (
	"fmt"
	"time"
)

func UpdateProductStock() <-chan map[string]int {
	stockUpdates := make(chan map[string]int)

	go func() {
		// add
		defer close(stockUpdates)

		currentStock := map[string]int{
			"Apples":  50,
			"Bananas": 30,
			"Oranges": 20,
			"Fr":      15,
		}

		for i := 0; i < 5; i++ {
			// add
			newStock := make(map[string]int)

			for product, quantity := range currentStock {
				newStock[product] = int(float64(quantity) * 0.95)
			}

			// add
			currentStock = newStock
			stockUpdates <- newStock

			time.Sleep(1 * time.Second)
		}
	}()

	return stockUpdates
}

func main() {
	stockStream := UpdateProductStock()

	var stockHistory []map[string]int

	/*old
	for i := 1; i < 5; i++ {
		stock := <- stockStream
		stockHistory = append(stockHistory, stock)
	}*/

	for stock := range stockStream {
		stockHistory = append(stockHistory, stock)
	}

	for i, stock := range stockHistory {
		fmt.Printf("Iteration %d: %v\n", i+1, stock)
	}
}

/*old3
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
}()
wg.Wait()
*/
