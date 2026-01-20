/*
fmt.Sprintf("%d", i)
оно не эффективно, так как внутри рефлексия
Рефлексия - дорогое определение типов

*/

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	//str := ""
	str2 := strings.Builder{}

	//println(reflect.TypeOf(str).String())

	start := time.Now()
	for i := 0; i < 100_000; i++ {
		//str2.WriteString(fmt.Sprintf("%d", i))
		str2.WriteString(strconv.Itoa(i))
		//str += fmt.Sprintf("%d", i)
	}
	fmt.Println(time.Since(start))
}
