package main

import (
	"fmt"
	"sort"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func main() {
	arrInt := []int{3, 2, 1, 5, 2, 4}                                // Создаём исходный массив int.
	arrFloat := []float64{3.33, 2.0, 0.3, 53.033, 2.0, 4.33}         // Создаём исходный массив float64.
	arrString := []string{"3", "three", "two", "one", "five", "two"} // Создаём исходный массив string.

	fmt.Println(sortInt(arrInt))
	fmt.Println(sortFloat(arrFloat))
	fmt.Println(sortString(arrString))
}

func sortInt(data []int) []int {
	// Проверяем отсортирован ли массив.
	ok := sort.IntsAreSorted(data)
	// Если нет, то сортиоруем.
	if !ok {
		sort.Ints(data)
	}

	return data // Возвращаем отсортированный массив.
}

func sortFloat(data []float64) []float64 {
	// Проверяем отсортирован ли массив.
	ok := sort.Float64sAreSorted(data)
	// Если нет, то сортиоруем.
	if !ok {
		sort.Float64s(data)
	}

	return data // Возвращаем отсортированный массив.
}

func sortString(data []string) []string {
	// Проверяем отсортирован ли массив.
	ok := sort.StringsAreSorted(data)
	// Если нет, то сортиоруем.
	if !ok {
		sort.Strings(data)
	}

	return data // Возвращаем отсортированный массив.
}
