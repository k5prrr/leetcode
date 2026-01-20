/*
Если смайл большой, и не помещяется в инт32
то он становится комбинированным
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "G🧑‍💻o"
	fmt.Println("Len: ", len(str))
	fmt.Println(str[1])
	fmt.Println(reflect.TypeOf(str[1]))
	//os.Exit(1)
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d %c \n", i, str[i])
	}
	fmt.Println("Next")
	for i, r := range str { // Когда рендч по строкам то бежим по рунамы
		fmt.Printf("%d %s %d %c \n", i, r, r, r)
	}
}
