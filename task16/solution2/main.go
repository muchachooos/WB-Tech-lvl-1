package main

import (
	"fmt"
	"sort"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func main() {
	fmt.Println(sortInt())
	fmt.Println(sortFloat())
	fmt.Println(sortString())
}

func sortInt() []int {
	// Создаём массив int и сортируем.
	return sort.IntSlice{3, 2, 1, 5, 2, 4}
}

func sortFloat() []float64 {
	// Создаём массив float64 и сортируем.
	return sort.Float64Slice{3.33, 2.0, 0.3, 53.033, 2.0, 4.33}
}

func sortString() []string {
	// Создаём массив string и сортируем.
	return sort.StringSlice{"3", "three", "two", "one", "five", "two"}
}
