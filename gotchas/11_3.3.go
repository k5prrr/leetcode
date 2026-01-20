/*

do modifyElement [10 20 30]
posle modifyElement [10 999 30]
do addElement [10 999 30]
slice in add:  [888 999 30 100]
posle addElement [10 999 30]



*/
package main

import (
	"fmt"
)


func modifyElement(slice []int) {
	slice[1] = 999
}
func addElement(slice []int) {
	slice = append(slice, 100)
	slice[0] = 888
	fmt.Println("slice in add: ", slice) // 888 999 30 100
}

func main() {
	original := []int{10, 20, 30}

	fmt.Println("do modifyElement", original)
	modifyElement(original)
	fmt.Println("posle modifyElement", original) // 10 999 30

	fmt.Println("do addElement", original)
	addElement(original)
	fmt.Println("posle addElement", original) // 10 999 30
}
