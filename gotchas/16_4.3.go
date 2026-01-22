/* Сервис для подсчёта уникальных слов
слов много и вы решили ограничить максимальное количество слов в мапе
если в мапу добавляется больше слов, чем лимит, она должна удалять старые записи
*/


package main

import (
	"fmt"
)


type WordCounter struct {
	counts map[string]int
	limit int
	
	words []string
}


func NewWordCounter(limit int) *WordCounter {
	return &WordCounter {
		counts:make(map[string]int),
		limit: limit,
	}
}

func (wc *WordCounter) AddWord(word string) {
	wc.counts[word]++
	wc.words = append(wc.words, word)

	if (len(wc.counts) < wc.limit) {
		return
	}

	delete(wc.counts, wc.words[0])
	wc.words = wc.words[1:]
}


func main() {
	wc := NewWordCounter(3)

	words := []string{"w0", "w1", "w0", "w2", "w3", "w4", "w5"}
	for _, word := range words {
		wc.AddWord(word)
	}

	fmt.Println(wc.counts)
}