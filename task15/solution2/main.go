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
	return sort.IntSlice{3, 2, 1, 5, 2, 4} // Создаём массив int и сортируем.
}

func sortFloat() []float64 {
	return sort.Float64Slice{3.33, 2.0, 0.3, 53.033, 2.0, 4.33} // Создаём массив float64 и сортируем.
}

func sortString() []string {
	return sort.StringSlice{"3", "three", "two", "one", "five", "two"} // Создаём массив string и сортируем.
}
