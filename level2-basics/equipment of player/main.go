package main

import (
	"fmt"
	"strings"
)

// Item представляет предмет в игре
type Item struct {
	Name  string
	Price int
}

// Inventory управляет коллекцией предметов
type Inventory struct {
	// Слайс для хранения порядка предметов
	items []Item
	// Мапа для мгновенного поиска по имени (O(1))
	cache map[string]int // Имя -> Индекс в слайсе
}

// NewInventory — конструктор для инициализации мапы
func NewInventory() *Inventory {
	return &Inventory{
		items: []Item{},
		cache: make(map[string]int),
	}
}

// AddItem добавляет предмет в инвентарь
func (inv *Inventory) AddItem(name string, price int) {
	newItem := Item{Name: name, Price: price}
	inv.items = append(inv.items, newItem)
	
	// Сохраняем индекс последнего добавленного предмета в мапу
	inv.cache[strings.ToLower(name)] = len(inv.items) - 1
	fmt.Printf("✅ Предмет '%s' добавлен.\n", name)
}

// FindItem ищет предмет по имени через мапу
func (inv *Inventory) FindItem(name string) (Item, bool) {
	index, safe := inv.cache[strings.ToLower(name)]
	if !safe {
		return Item{}, false
	}
	return inv.items[index], true
}

// RemoveItem удаляет предмет (простой способ через пересоздание мапы)
func (inv *Inventory) RemoveItem(name string) {
	key := strings.ToLower(name)
	index, exists := inv.cache[key]
	if !exists {
		fmt.Printf("❌ Предмет '%s' не найден.\n", name)
		return
	}

	// Удаляем из слайса (быстрый способ без сохранения порядка)
	// Меняем удаляемый элемент с последним и отрезаем хвост
	lastIdx := len(inv.items) - 1
	inv.items[index] = inv.items[lastIdx]
	
	// Обновляем индекс перемещенного элемента в кэше
	inv.cache[strings.ToLower(inv.items[index].Name)] = index
	
	// Отрезаем последний элемент
	inv.items = inv.items[:lastIdx]
	delete(inv.cache, key)
	
	fmt.Printf("🗑️ Предмет '%s' удален.\n", name)
}

// Show выводит весь инвентарь
func (inv *Inventory) Show() {
	fmt.Println("\n--- Содержимое инвентаря ---")
	if len(inv.items) == 0 {
		fmt.Println("Пусто")
		return
	}
	for i, item := range inv.items {
		fmt.Printf("%d. %s (%d золота)\n", i+1, item.Name, item.Price)
	}
}

func main() {
	inv := NewInventory()

	// Добавляем предметы
	inv.AddItem("Меч Гофера", 150)
	inv.AddItem("Зелье маны", 50)
	inv.AddItem("Щит из кода", 300)

	inv.Show()

	// Поиск
	search := "Зелье маны"
	if item, ok := inv.FindItem(search); ok {
		fmt.Printf("\n🔍 Найдено: %s за %d монет.\n", item.Name, item.Price)
	}

	// Удаление
	inv.RemoveItem("Зелье маны")
	inv.Show()
}
