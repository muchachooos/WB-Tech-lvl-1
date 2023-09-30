package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные.
// (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

func main() {
	// Создаём строку.
	message := "abvgEde"

	fmt.Println(message, uniqueness(message))
}

func uniqueness(msg string) bool {
	// Приводим msg к нижнему регистру.
	msgLow := strings.ToLower(msg)

	// Создаём слайс, в который будем сохранять символы для сравнения.
	symbols := make([]byte, 0, len(msgLow))
	// Проходим циклом по msgLow.
	for i := range msgLow {
		// (Тут можно использовать slices.Contains.)
		if contains(symbols, msgLow[i]) {
			return false // Если символ уже содержится в слайсе symbols, то возвращается значение false.
		}

		// Если такого символа в symbols нет, то добавляем его.
		symbols = append(symbols, msgLow[i])
	}

	return true
}

// Ф-ция contains проверяет слайс s на наличие эл-та v.
func contains(s []byte, v byte) bool {
	// Проходим циклом по s.
	for i := range s {
		if s[i] == v { // Если элемент равен v, то возвращаем true.
			return true
		}
	}

	// Если совпадений нет, возвращаем false.
	return false
}
