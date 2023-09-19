package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	var a, b int

	fmt.Print("Введите число a: ")
	fmt.Scan(&a)
	fmt.Print("Введите число b: ")
	fmt.Scan(&b)

	a, b = b, a

	fmt.Printf("a=%d, b=%d", a, b)
}
