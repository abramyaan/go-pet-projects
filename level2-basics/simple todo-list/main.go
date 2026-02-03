package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Task представляет одну задачу
type Task struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Done  bool   `json:"done"`
}

// TodoList управляет списком задач
type TodoList struct {
	Tasks    []Task `json:"tasks"`
	LastID   int    `json:"last_id"`
	Filename string `json:"-"`
}

// NewTodoList создает новый список и загружает данные из файла
func NewTodoList(filename string) *TodoList {
	list := &TodoList{
		Tasks:    []Task{},
		Filename: filename,
	}
	list.load()
	return list
}

// Add добавляет новую задачу
func (l *TodoList) Add(text string) {
	l.LastID++
	newTask := Task{
		ID:   l.LastID,
		Text: text,
		Done: false,
	}
	l.Tasks = append(l.Tasks, newTask)
	l.save()
	fmt.Println("✅ Задача добавлена!")
}

// List выводит все задачи
func (l *TodoList) List() {
	fmt.Println("\n--- СПИСОК ЗАДАЧ ---")
	if len(l.Tasks) == 0 {
		fmt.Println("Список пуст.")
		return
	}
	for _, task := range l.Tasks {
		status := " "
		if task.Done {
			status = "X"
		}
		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Text)
	}
}

// Done отмечает задачу как выполненную
func (l *TodoList) Done
