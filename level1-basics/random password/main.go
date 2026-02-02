package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
	symbols      = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

func generatePassword(length int, useUpper, useDigits, useSymbols bool) string {
	// Инициализируем генератор случайных чисел текущим временем
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Создаем набор доступных символов на основе настроек
	charset := lowerLetters
	if useUpper {
		charset += upperLetters
	}
	if useDigits {
		charset += digits
	}
	if useSymbols {
		charset += symbols
	}

	// Используем strings.Builder для эффективного построения строки
	var password strings.Builder
	// Предварительно выделяем память (оптимизация)
	password.Grow(length)

	for i := 0; i < length; i++ {
		// Выбираем случайный индекс из набора символов
		randomIndex := r.Intn(len(charset))
		// Записываем байт в буфер
		password.WriteByte(charset[randomIndex])
	}

	return password.String()
}

func main() {
	// Настройки пароля
	length := 16
	includeUpper := true
	includeDigits := true
	includeSymbols := true

	fmt.Println("--- Генератор паролей ---")
	
	// Генерируем несколько вариантов для примера
	for i := 1; i <= 5; i++ {
		pass := generatePassword(length, includeUpper, includeDigits, includeSymbols)
		fmt.Printf("Пароль %d: %s\n", i, pass)
	}
}
