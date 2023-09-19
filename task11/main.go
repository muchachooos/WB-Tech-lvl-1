package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func main() {
	s1 := []int{1, 2, 3, 3, 4, 5} // первое множество.
	s2 := []int{3, 4, 5, 5, 6, 7} // второе множество.

	fmt.Println(intersection(s1, s2))

}

func intersection(s1, s2 []int) []int {
	inter := make([]int, 0) // создаем пустой массив для хранения пересечения.

	for _, val1 := range s1 { // перебираем элементы первого множества.
		for _, val2 := range s2 { // перебираем элементы второго множества.
			if val1 == val2 { // если элементы равны добавляем его в массив пересечения
				inter = append(inter, val1)
			}
		}
	}

	return inter
}
