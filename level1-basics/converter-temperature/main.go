package main

import (
	"fmt"
)


func fromCelsius(val float64, target string) float64 {
	switch target {
	case "F":
		return (val * 9 / 5) + 32
	case "K":
		return val + 273.15
	default:
		return val
	}
}


func toCelsius(val float64, source string) float64 {
	switch source {
	case "F":
		return (val - 32) * 5 / 9
	case "K":
		return val - 273.15
	default:
		return val
	}
}

func main() {
	var val float64
	var choice int

	fmt.Println("=== Конвертер температуры ===")
	fmt.Println("1. Цельсий (C) -> Фаренгейт (F)")
	fmt.Println("2. Фаренгейт (F) -> Цельсий (C)")
	fmt.Println("3. Цельсий (C) -> Кельвин (K)")
	fmt.Println("4. Кельвин (K) -> Цельсий (C)")
	fmt.Print("\nВыберите вариант (1-4): ")
	fmt.Scan(&choice)

	fmt.Print("Введите значение температуры: ")
	fmt.Scan(&val)

	var result float64
	var unit string

	switch choice {
	case 1:
		result = fromCelsius(val, "F")
		unit = "°F"
	case 2:
		result = toCelsius(val, "F")
		unit = "°C"
	case 3:
		result = fromCelsius(val, "K")
		unit = "K"
	case 4:
		result = toCelsius(val, "K")
		unit = "°C"
	default:
		fmt.Println("Ошибка: неверный выбор.")
		return
	}

	// Выводим результат с округлением до 2 знаков после запятой
	fmt.Printf("\nРезультат: %.2f %s\n", result, unit)
}
