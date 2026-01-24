package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	vals = 50
	maxInt = 9
)

func main() {
	//var ch1 chan int
	ch1 := make(chan int, 0)
	ch2 := make(chan int, 16)

	go func() {
		for i := 0; i < vals; i++ {
			if  i == 25 {
				fmt.Print("\neee 25\n")
				close(ch1)
			}

			r:= rand.Intn(maxInt)
			fmt.Printf("%d ", r)

			if i < 25 {
				ch1 <- r
				continue
			}

			ch2 <- r
		}
		close(ch2)
	}()

	go func() {
		for vl := range ch1{
			fmt.Printf("\n1i [%d]\n", vl)
		}
		for vl := range ch2{
			fmt.Printf("\n2i [%d]\n", vl)
		}
	}()

	time.Sleep(time.Second*7)
	fmt.Println("\nend")
}
