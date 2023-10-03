// Определяем главную функцию main
package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20.

func main() {
	a := *big.NewInt(1234567890)
	b := *big.NewInt(9876543210)

	// Создаём переменные для вычесления минимального значения.
	var x int64 = 2
	var deg int64 = 20

	// Объявляем переменную minValue типа *big.Int со значением 2^20.
	minValue := new(big.Int).Exp(big.NewInt(x), big.NewInt(deg), nil)

	// Если значения меньше чем минимальное - паника.
	if a.Cmp(minValue) < 0 && b.Cmp(minValue) < 0 {
		panic("Please enter numbers greater than 2^20")
	}

	// Если больше,то проводим вычесления.
	mulRes := new(big.Int).Mul(&a, &b) // Умножаем.
	divRes := new(big.Int).Div(&a, &b) // Делим.
	addRes := new(big.Int).Add(&a, &b) // Складывае.
	subRes := new(big.Int).Sub(&a, &b) // Вычитаем.

	// Выводим результаты.
	fmt.Println("Multiplication result:", mulRes)
	fmt.Println("Division result:", divRes)
	fmt.Println("Addition result:", addRes)
	fmt.Println("Subtraction result:", subRes)
}
