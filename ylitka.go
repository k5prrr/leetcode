package main

import (
	"bytes"
	"fmt"
)

type Ylitka struct {
	table  [][]int
	width  int
	height int
}

func NewYlitka(width, height int) *Ylitka {
	return &Ylitka{
		width:  width,
		height: height,
	}
}
func PlayYlitka(width, height int) {
	yl := NewYlitka(width, height)
	yl.CreateTable()
	yl.Fill()
	yl.Show()
}
func (y *Ylitka) CreateTable() {
	y.table = make([][]int, y.height)
	for i := range y.table {
		y.table[i] = make([]int, y.width)
	}
}
func (y *Ylitka) Show() {
	var b bytes.Buffer
	for _, tableY := range y.table {
		for _, value := range tableY {
			b.WriteString(fmt.Sprintf("%d\t", value))
		}
		b.WriteString("\n")
	}
	fmt.Println(b.String())
}
func (y *Ylitka) Fill() {
	top := 0
	bottom := y.height - 1
	left := 0
	right := y.width - 1

	k := 0
	for left <= right && top <= bottom {
		// ->
		// 00
		for i := left; i <= right; i++ {
			k++
			y.table[top][i] = k
		}
		top++

		// 0|
		// 0V
		for i := top; i <= bottom; i++ {
			k++
			y.table[i][right] = k
		}
		right--

		// exit
		if left > right || top > bottom {
			break
		}

		// 00
		// <-
		for i := right; i >= left; i-- {
			k++
			y.table[bottom][i] = k
		}
		bottom--

		// ^0
		// |0
		for i := bottom; i >= top; i-- {
			k++
			y.table[i][left] = k
		}
		left++
	}
}
func main() {
	PlayYlitka(5, 7)
	PlayYlitka(7, 5)
}
