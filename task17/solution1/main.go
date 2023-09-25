package main

import (
	"fmt"
	"sort"
)

// Реализовать бинарный поиск встроенными методами языка.

func main() {
	// Создаем три слайса разных типов данных.
	arrInt := []int{1, 3, 4, 5}
	arrString := []string{"3", "five", "one", "three"}
	arrFloat := []float64{0.33, 3, 3.33, 12}

	// Используем ф-ции поиска SearchType.
	// Передаём слайс и значение которое ищем.
	fmt.Println(sort.SearchInts(arrInt, 3))         // Ф-ция возвращает индекс искомого значения.
	fmt.Println(sort.SearchStrings(arrString, "3")) // Ф-ция возвращает индекс искомого значения.
	fmt.Println(sort.SearchFloat64s(arrFloat, 32))  // Или индекс для его вставки, если оно отсутствует.
}
