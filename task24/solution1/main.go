package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point
// с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x float64
	y float64
}

func main() {
	// Создаем две точки.
	p1 := NewPoint(-1, 3)
	p2 := NewPoint(2, 4)

	// Вычисляем расстояние между точками.
	distance := distance(p1, p2)

	// Выводим результат.
	fmt.Println(distance)
}

// DistanceTo вычисляет расстояние между текущей точкой и другой заданной точкой.
func distance(pA, pB *Point) float64 {
	// Вычисдяем разницу между координатами по осям, что будет являться катетами.
	legX := pA.x - pB.x
	legY := pA.y - pB.y

	// Вычисляем квадратный корень из суммы квадвратов катетов и получаем расстояние между точками.
	res := math.Sqrt(legX*legX + legY*legY)

	// Возвращаем результат.
	return res
}

// NewPoint создает новую точку с заданными координатами и возвращает указатель на нее.
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}
