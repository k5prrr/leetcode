package main
// Замыкание, когда внутри функции используем переменную, обьявленую снаружи
import "fmt"

func unID() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

var k func() int = unID()

func main() {
	fmt.Println(k())
	fmt.Println(k())
	fmt.Println(k())
}
