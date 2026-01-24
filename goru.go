package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 9; i++ {
		go func(i int) {
			fmt.Println(i)
			time.Sleep(time.Second)
		}(i)
	}
	time.Sleep(time.Second * 5)
}
