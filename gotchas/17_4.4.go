/*
Есть мапа, кде ключ строка, а значение слайс строк
каждый слайс содержит набор уникальных значений для ключа
надо написать функцию MergeToMap, которая принимает
мапу и новый слайс для конкретного ключа,
анализирует существующие и добавляет только то, что там ещё нет

пример
m := map[string][]string{
	"g1": {"a", "b"},
	"g2": {"c"},
}

newKey := "g1"
mewSlice := []string{"a" , "d"}

Результат
m := map[string][]string{
	"g1": {"a", "b", "d"},
	"g2": {"c"},
}


*/

package main

import "fmt"

func MergeToMap(m *map[string][]string, key string, valueSlice []string) {
	uniqExists := make(map[string]struct{})
	for _, v := range (*m)[key] {
		uniqExists[v] = struct{}{}
	}

	for _, v := range valueSlice {
		if _, ok := uniqExists[v]; !ok {
			(*m)[key] = append((*m)[key], v)
			uniqExists[v] = struct{}{}
		}
	}
}

func MergeToMap0(m *map[string][]string, key string, valueSlice []string) {

	// Уникально что пришло
	uniqValueSliceMap := make(map[string]struct{})
	var uniqValueSlice []string
	for _, word := range valueSlice {
		if _, ok := uniqValueSliceMap[word]; ok {
			continue
		}
		uniqValueSliceMap[word] = struct{}{}
		uniqValueSlice = append(uniqValueSlice, word)
	}

	// Создание подобного
	slice, ok := (*m)[key]
	if !ok {
		(*m)[key] = uniqValueSlice
		return
	}

	tmp := make(map[string]struct{})
	for _, word := range slice {
		tmp[word] = struct{}{}
	}

	for _, word := range uniqValueSlice {
		if _, ok := tmp[word]; ok {
			continue
		}
		(*m)[key] = append((*m)[key], word)
	}

}

func main() {
	m := map[string][]string{
		"g1": {"a", "b"},
		"g2": {"c"},
	}
	MergeToMap(&m, "g1", []string{"b", "d"})
	MergeToMap(&m, "g3", []string{"b", "d"})
	MergeToMap(&m, "g4", []string{"b", "b", "b", "d"})
	fmt.Println(m)
}
