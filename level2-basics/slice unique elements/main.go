package main

import (
	"fmt"
)

// Unique возвращает новый слайс, содержащий только уникальные элементы.
func Unique(input []string) []string {
	// Создаем мапу, которая будет работать как "множество" (Set).
	// Используем struct{} в качестве значения, так как она занимает 0 байт.
	seen := make(map[string]struct{})
	
	// Заранее выделяем емкость для результата, чтобы избежать лишних аллокаций.
	// В худшем случае (все уникальны) размер будет len(input).
	result := make([]string, 0, len(input))

	for _, val := range input {
		// Проверяем, встречали ли мы этот элемент ранее
		if _, ok := seen[val]; !ok {
			// Если нет — фиксируем его в мапе и добавляем в результат
			seen[val] = struct{}{}
			result = append(result, val)
		}
	}

	return result
}

func main() {
	// Исходный слайс с дубликатами
	words := []string{"go", "rust", "go", "python", "rust", "cpp", "go"}

	fmt.Println("Исходный список:", words)
	
	uniqueWords := Unique(words)
	
	fmt.Println("Уникальные элементы:", uniqueWords)
	fmt.Printf("Длина: %d, Емкость: %d\n", len(uniqueWords), cap(uniqueWords))
}
