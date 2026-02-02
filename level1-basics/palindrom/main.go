package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// IsPalindrome проверяет, является ли строка палиндромом.
// Использует срез рун для корректной работы с UTF-8 (кириллицей).
func IsPalindrome(s string) bool {
	// Преобразуем строку в срез рун, чтобы правильно обрабатывать многобайтовые символы
	runes := []rune{}
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			runes = append(runes, unicode.ToLower(r))
		}
	}

	// Алгоритм двух указателей
	left := 0
	right := len(runes) - 1

	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// IsNumberPalindrome проверяет число без конвертации в строку (математически)
func IsNumberPalindrome(n int) bool {
	if n < 0 {
		return false
	}

	original := n
	reversed := 0

	for n > 0 {
		remainder := n % 10
		reversed = reversed*10 + remainder
		n /= 10
	}

	return original == reversed
}

func main() {
	// Тесты для строк
	words := []string{"А роза упала на лапу Азора
