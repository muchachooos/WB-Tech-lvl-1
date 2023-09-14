package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	var a, b int

	fmt.Print("Введите число a: ")
	fmt.Scan(&a)
	fmt.Print("Введите число b: ")
	fmt.Scan(&b)

	// a = 10
	// b = 20
	a = a + b // a = 10 + 20 = 30
	b = a - b // b = 30 - 20 = 10
	a = a - b // a = 30 - 10 = 20

	fmt.Printf("a=%d, b=%d", a, b)
}
