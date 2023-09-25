package main

import (
	"fmt"
	"sort"
)

// Реализовать бинарный поиск встроенными методами языка.

func main() {
	//Создаем отсортированный массив целых чисел.
	arr := []int{2, 5, 8, 10, 13, 16, 20}

	// Будем искать значение 13.
	target := 13

	// Используем функцию Search для поиска индекса элемента.
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	if index < len(arr) && arr[index] == target { // Если элемент найден, выводим значение и его индекс.
		fmt.Printf("Found %d at index %d\n", target, index)
	} else {
		fmt.Printf("%d not found in array\n", target)
	}
}
