package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

// Pair — вспомогательная структура для сортировки мапы
type Pair struct {
	Key   string
	Value int
}

func main() {
	text := `Go is an open source programming language that makes it easy to build 
	simple, reliable, and efficient software. Go is efficient, Go is simple!`

	// 1. Очистка и разбиение текста на слова
	// Убираем пунктуацию и приводим к нижнему регистру
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	words := strings.FieldsFunc(strings.ToLower(text), f)

	// 2. Подсчет частоты слов через map
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}

	// 3. Сортировка мапы по значению (частоте)
	// В Go нельзя напрямую сортировать мапу, поэтому переносим данные в слайс структур
	var sortedPairs []Pair
	for k, v := range counts {
		sortedPairs = append(sortedPairs, Pair{k, v})
	}

	// Сортируем: сначала по количеству (убывание), потом по алфавиту
	sort.Slice(sortedPairs, func(i, j int) bool {
		if sortedPairs[i].Value != sortedPairs[j].Value {
			return sortedPairs[i].Value > sortedPairs[j].Value
		}
		return sortedPairs[i].Key < sortedPairs[j].Key
	})

	// 4. Вывод результата
	fmt.Printf("%-12s | %s\n", "Слово", "Частота")
	fmt.Println("-------------------------")
	for _, p := range sortedPairs {
		fmt.Printf("%-12s | %d\n", p.Key, p.Value)
	}
}
