package main

import (
	"fmt"
	"math/rand"
)

func main() {
	secretNumber := rand.Intn(100) + 1 
	attempts := 0
	var guess int
	fmt.Println("Я загадал число от 1 до 100. Попробуй угадать!")

	for {
		fmt.Print("Твоё предположение: ")
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Пожалуйста, вводи только числа!")
			continue
		}

		attempts++

		if guess == secretNumber {
			fmt.Printf("Угадал! Поздравляю! Это было число %d.\n", secretNumber)
			fmt.Printf("Ты справился за %d попыток.\n", attempts)
			break
		} else if guess < secretNumber {
			fmt.Println("Слишком мало! Попробуй больше.")
		} else {
			fmt.Println("Слишком много! Попробуй меньше.")
		}
	}

	fmt.Println("Игра окончена. Хочешь сыграть ещё? (y/n)")
	var answer string
	fmt.Scan(&answer)
	if answer == "y" || answer == "Y" {
		main() 
	}
}
