package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func main() {
	msg := "snow dog sun"

	// Разбиваем строку msg на слова.
	words := strings.Split(msg, " ")

	// Создаём слайс для записи результата.
	result := make([]string, len(words))

	// Перебираем все слова в words и записываем их в result с конца.
	for i, val := range words {
		result[len(words)-i-1] = val
	}

	// Собираем слова обратно в строку и выводим.
	fmt.Println(strings.Join(result, " "))
}
