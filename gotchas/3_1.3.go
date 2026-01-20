package main

import (
	"fmt"
)

type User struct {
	Balance int
}

func main() {
	balance := 1000
	user := &User{balance}

	defer printBalance("Старт баланс", user.Balance)          //1000
	defer printBalance("Текущий баланс", user.Balance)        //1000
	defer printUserBalance("Указатель на пользователя", user) //1300 так как он запоминает указатель на старый объект

	user.Balance += 500           //1500
	updateUserBalance(user, -200) //1300

	user = &User{Balance: 300} // 300
}
func updateUserBalance(u *User, kol int) {
	u.Balance += kol
}
func printBalance(text string, balance int) {
	fmt.Printf("%s = %d\n", text, balance)
}
func printUserBalance(text string, u *User) {
	fmt.Println(u)
	fmt.Printf("%s = %d\n", text, u.Balance)
}

/*
&{1300}
Указатель на баланс = 1300
Текущий баланс = 1000
Старт баланс = 1000


*/
