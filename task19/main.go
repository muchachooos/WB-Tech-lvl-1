package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку.

func main() {
	fmt.Println(reverse("Hello"))
}

func reverse(msg string) string {
	var res string

	// Проходимся циклом по всем символам строки msg.
	for _, val := range msg {
		res = string(val) + res // На каждой итерации мы добавляем текущий символ в начало строки res.
	}

	return res
}
