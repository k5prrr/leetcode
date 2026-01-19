package main

import "fmt"

type User struct {
	Name string
}

func main() {
	user := User{Name: "Евгений"}
	fmt.Println("Имя до обновления:", user.Name) // Евгений

	updateUser(user)
	fmt.Println("Имя после обновления:", user.Name) // Евгений
}

func updateUser(u User) {
	u.Name = "Саня"
	fmt.Println("Имя внутри функции [updateUser]:", u.Name) // Саня

	resetUser(&u)
	fmt.Println("Имя после вызова функции [resetUser]:", u.Name) // NN
}

func resetUser(u *User) {
	u.Name = "NN"
	u = &User{Name: "Безымянный"}
	fmt.Println("Имя внутри функции [resetUser]:", u.Name) // Безымянный
}
/*
 * Имя до обновления: Евгений
 * Имя внутри функции [updateUser]: Саня
 * Имя внутри функции [resetUser]: Безымянный
 * Имя после вызова функции [resetUser]: NN
 * Имя после обновления: Евгений
 *
 *
 * Указатель примерно 8 байт весит
 *
 */
