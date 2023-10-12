package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.

func main() {
	// Создаем множество.
	set := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(set)

	// Создаём подмножество Set.
	subset := set[1:4]
	fmt.Println(subset)
}

// Собственное (истинное) подмножество — C является подмножеством A, но C != A;
// Каждое подмножество по определению тоже множество.

// https://habr.com/ru/articles/457312/
