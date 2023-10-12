package main

import "fmt"

type Mac struct {
	turnOn bool
}

// TurnOn Метод включающий Mac.
func (m *Mac) TurnOn() {
	// Включаем Mac.
	m.turnOn = true

	fmt.Println("Mac on")
}

// TurnOff Метод выключающий Mac.
func (m *Mac) TurnOff() {
	// Выключаем Mac.
	m.turnOn = false

	fmt.Println("Mac off")
}
