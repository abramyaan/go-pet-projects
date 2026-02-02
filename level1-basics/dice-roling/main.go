package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// RollDice имитирует бросок n кубиков с d гранями
func RollDice(count, sides int) []int {
	results := make([]int, count)
	for i := 0; i < count; i++ {
		// rand.Intn(6) дает 0-5, поэтому добавляем 1
		results[i] = rand.Intn(sides) + 1
	}
	return results
}

// Sum вычисляет сумму значений в слайсе
func Sum(results []int) int {
	total := 0
	for _, v := range results {
		total += v
	}
	return total
}

func main() {
	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	// 1. Бросок 2d6 (два шестигранных кубика)
	fmt.Println("--- Бросок 2d6 ---")
	d26 := RollDice(2, 6)
	fmt.Printf("Результаты: %v | Сумма: %d\n\n", d26, Sum(d26))

	// 2. Бросок 3d6 с сортировкой
	fmt.Println("--- Бросок 3d6 (с сортировкой) ---")
	d36 := RollDice(3, 6)
	sort.Ints(d36) // Сортируем слайс по возрастанию
	fmt.Printf("Отсортировано: %v | Сумма: %d\n\n", d36, Sum(d36))

	// 3. Механика преимущества (бросаем 2d20, берем большее)
	fmt.Println("--- Бросок d20 с преимуществом ---")
	adv := RollDice(2, 20)
	
	// Используем сортировку, чтобы легко найти максимум
	sort.Ints(adv)
	best := adv[1] // Последний элемент после сортировки - наибольший
	
	fmt.Printf("Броски: %v | Результат с преимуществом: %d\n", adv, best)
}
