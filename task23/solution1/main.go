package main

import "fmt"

// Удалить i-тый эллемент из слайса.

func main() {
	s := []int{12, 33, 1, 2, 33, 5}

	fmt.Println(delElement(s, 1))
}

func delElement(s []int, target int) []int {
	return append(s[:target], s[target+1:]...)
}
