package main

import (
	"fmt"
)

type SomeStruct struct {
	Value int
}

func CheckForNil(i interface{}) { // Если б исправил на *SomeStruct , то был бы nil
	if i == nil {
		fmt.Println("nil")
		return
	}

	fmt.Println("not nil")
}

func main() {
	var s *SomeStruct // s = nil, тип = *SomeStruct
	CheckForNil(s)    // передаётся в interface{}
}

/*
 * interface value:
 ├── тип: *SomeStruct     ← НЕ nil
 └── значение: nil        ← nil
 В Go interface{} == nil возвращает true
 только если и тип, и значение равны nil.
 В текущем случае тип не nil, поэтому сравнение возвращает false.
*/
