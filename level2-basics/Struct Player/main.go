package main

import (
	"fmt"
)

// Stats — отдельная структура для характеристик (композиция)
type Stats struct {
	Level      int
	Experience int
}

// Player — основная структура игрока
type Player struct {
	Name   string
	HP     int
	MaxHP  int
	Stats  // Встраивание (Embedding) для демонстрации композиции
}

// TakeDamage — Pointer Receiver (*). Изменяет состояние оригинального объекта.
func (p *Player) TakeDamage(amount int) {
	p.HP -= amount
	if p.HP < 0 {
		p.HP = 0
	}
	fmt.Printf("💥 %s получил %d урона. Текущее HP: %d\n", p.Name, amount, p.HP)
}

// Heal — Pointer Receiver (*). Позволяет восстановить здоровье.
func (p *Player) Heal(amount int) {
	p.HP += amount
	if p.HP > p.MaxHP {
		p.HP = p.MaxHP
	}
	fmt.Printf("❤️ %s вылечился на %d. Текущее HP: %d\n", p.Name, amount, p.HP)
}

// LevelUp — Pointer Receiver (*). Повышает уровень и увеличивает MaxHP.
func (p *Player) LevelUp() {
	p.Level++
	p.MaxHP += 20
	p.HP = p.MaxHP // Полное восстановление при повышении уровня
	fmt.Printf("🆙 УРОВЕНЬ ПОВЫШЕН! Теперь %s %d уровня. MaxHP: %d\n", p.Name, p.Level, p.MaxHP)
}

// GetStatus — Value Receiver (без *). Работает с копией, только для чтения.
func (p Player) GetStatus() string {
	return fmt.Sprintf("[%s] Уровень: %d, HP: %d/%d", p.Name, p.Level, p.HP, p.MaxHP)
}

func main() {
	// Инициализация игрока
	hero := Player{
		Name:  "GopherKnight",
		HP:    100,
		MaxHP: 100,
		Stats: Stats{
			Level:      1,
			Experience: 0,
		},
	}

	fmt.Println("--- Начало приключения ---")
	fmt.Println(hero.GetStatus())

	// Симулируем события
	hero.TakeDamage(30)
	hero.Heal(10)
	hero.LevelUp()
	
	fmt.Println("\n--- Финальный статус ---")
	fmt.Println(hero.GetStatus())
}
