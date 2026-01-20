package main

// Пробовал встраивание

import (
	"fmt"
)

type Animal struct {
	name string
	age string
}
func NewAnimal(name string) *Animal {
	return &Animal{name: name}
}

func (a *Animal) GetName() string {
	return a.name
}

func (a *Animal) SetName(name string) {
	a.name = name
}


type Cat struct {
	Animal
	color  string
}


func NewCat(name, color string) *Cat {
	return &Cat{
		Animal: Animal{name: name},
		color:  color,
	}
}


func (c *Cat) GetColor() string {
	return c.color
}

func main() {
	cat := NewCat("Whiskers", "black")
	fmt.Println("Name:", cat.GetName())
	fmt.Println("Color:", cat.GetColor())

	cat.SetName("Tom")
	fmt.Println("New name:", cat.GetName())
	fmt.Sprintf("%v", cat)
}
