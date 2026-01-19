package main

import (
	"fmt"
	"strconv"
)

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль запрещено")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("неизвестная операция: %s (допустимы: +, -, *, /)", op)
	}
}

func main() {
	var inputA, inputB, operation string

	fmt.Println("Простой калькулятор (+, -, *, /)")
	fmt.Println("Введите первое число:")
	fmt.Scan(&inputA)

	fmt.Println("Введите второе число:")
	fmt.Scan(&inputB)

	fmt.Println("Введите операцию (+, -, *, /):")
	fmt.Scan(&operation)

	// Преобразуем строки в float64
	a, errA := strconv.ParseFloat(inputA, 64)
	b, errB := strconv.ParseFloat(inputB, 64)

	if errA != nil || errB != nil {
		fmt.Println("Ошибка: введите корректные числа")
		return
	}

	result, err := calculate(a, b, operation)

	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Результат: %.4f\n", result)
	}
}
