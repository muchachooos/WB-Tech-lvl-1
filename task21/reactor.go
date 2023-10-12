package main

import "fmt"

// NuclearReactor Тип реактора.
type NuclearReactor struct {
	turnOn bool
}

// Power Метод NuclearReactor который изменяет его состояние.
func (n *NuclearReactor) Power(turnOn bool) {
	// Переводим реактор в состояние turnOn.
	n.turnOn = turnOn

	// Выводим состояние реактора.
	fmt.Println("Reactor condition: ", n.turnOn)
}
