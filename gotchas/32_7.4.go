/* пайплайн
вы разрабатываете систему для обработки финансовых транзакций
каждая транзакция проходит несколько этапов обработки
1. чтение транзакций из исходных данных
2. фильтрация транзакций: убираем всё с отрицательными суммами
3. Конвертация валюты: преобразование в доллары
4. Сохранение результатов

просто знай, что на все стадии пайплайна лучше пробрасывать контекст, чтоб их контролировать
 */
package main

import (
	"fmt"
	"math/rand"
)

type Transaction struct {
	ID int64
	Amount int64 // Хранение в копейках
}
func filterTransactions(in <- chan Transaction) <-chan Transaction {
	out := make(chan Transaction)

	go func() {
		for tr := range in {
			if tr.Amount >= 0 {
				out <- tr
			}
		}

		close(out)
	}()

	return out
}

func convertTransactions(in <- chan Transaction) <-chan Transaction {
	out := make(chan Transaction)

	go func() {
		for tr := range in {
			tr.Amount = int64(float64(tr.Amount) * 0.8)
			out <- tr
		}

		close(out)
	}()

	return out
}
func generateTransactions(count int) <-chan Transaction {
	out := make(chan Transaction)

	go func() {
		for i := 0; i< count; i++ {
			out <- Transaction{
				ID: int64(i),
				Amount: rand.Int63(),
			}
		}

		close(out)
	}()

	return out
}
func showTransactions(in <- chan Transaction) {

	for tr := range in {
		fmt.Printf("Transaction ID: %d, Amount: %d\n", tr.ID, tr.Amount)
	}
	fmt.Println("end")
}

func main() {
	transactions := generateTransactions(10)
	filtered := filterTransactions(transactions)
	converted := convertTransactions(filtered)
	showTransactions(converted)

}
