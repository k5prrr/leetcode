/*
defer при выходе из телаИлиФункции,
В обратном порядке (Со стека)

Панику можно перехватить в defer recaver
Но лог фатал, это вывод лога и вызов os.exit
*/
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	x := 10
	y := 20

	fmt.Println(reflect.TypeOf(x), unsafe.Sizeof(x))

	// Тут он уже сохранился, возьмёт с памяти 10
	defer func(val int) {
		fmt.Println("x:", val)
	}(x)

	// Тут он возьмёт с глобальной видимости
	defer func() {
		fmt.Println("y:", y) // Замыкание (так как тут на указатели)
		// Замыкание, когда переменная создана за областью видимости(тела)
	}()

	x = 100
	y = 200

	fmt.Println("END")
}

/*
int 8
END
y: 200
x: 10

*/
