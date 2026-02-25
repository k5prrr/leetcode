/* пайплайн
вы разрабатываете систему для обработки финансовых транзакций
каждая транзакция проходит несколько этапов обработки
1. чтение транзакций из исходных данных
2. фильтрация транзакций: убираем всё с отрицательными суммами
3. Конвертация валюты: преобразование в доллары
4. Сохранение результатов
 */
package main

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
			tr.Amount *= 0.8
			out <- tr
		}

		close(out)
	}()

	return out
}

func saveTransactions(in <- chan Transaction) {
	go func() {
		for tr := range in {
			fmt.Printf("Transaction ID: %d, Amount: %.2\n", tr.ID, tr.Amount)
		}
	}()
}

func main() {}
