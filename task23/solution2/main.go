package main

import "fmt"

// Удалить i-тый эллемент из слайса.

func main() {
	s := []int{12, 33, 1, 2, 33, 5}

	fmt.Println(delElement(s, 1))
}

func delElement(s []int, target int) []int {
	// Создаем новый слайс с размером на 1 меньше исходного
	result := make([]int, len(s)-1)
	// Копируем элементы до i-го элемента
	copy(result, s[:target])
	// Копируем элементы после i-го элемента
	copy(result[target:], s[target+1:])

	return result
}
