package main

import "fmt"

type Car struct {
	color   string
	mileage int
}

func main() {
	cars := []Car{
		Car{color: "red", mileage: 10},
		Car{color: "blue", mileage: 20},
		Car{color: "green", mileage: 3_0},
	}
	fmt.Printf("%d %d %p %p \n", len(cars), cap(cars), &cars, &cars[0])
	carN := &cars[0]
	carN.mileage += 100

	cars = append(cars, Car{color: "red2", mileage: 1000}) // Это перемещает слайс
	fmt.Printf("%d %d %p %p \n", len(cars), cap(cars), &cars, &cars[0])
	carN.mileage += 100 // А этот плюсует к старому месту

	fmt.Println(cars[0].mileage, cars[0].color) // 110
	fmt.Println(carN.mileage, carN.color)       // 210
}

/*
У слайса одинаковый, так как он меняется структурой
3 3 0xc00000e018 0xc000110000
4 6 0xc00000e018 0xc000116000
110 red
210 red
*/
