package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func main() {
	s1 := []int{1, 2, 3, 3, 4, 5} // первое множество.
	s2 := []int{3, 4, 5, 5, 7}    // второе множество.

	fmt.Println(intersection(s1, s2))

}

func intersection(s1, s2 []int) []int {
	var length int

	// Сохраняем длину меньшего массива в переменную length.
	if len(s1) < len(s2) {
		length = len(s1)
	} else {
		length = len(s2)
	}

	// создаем пустой массив для хранения пересечения.
	inter := make([]int, 0, length)

	for _, val1 := range s1 { // перебираем элементы первого множества.
		for _, val2 := range s2 { // перебираем элементы второго множества.
			if val1 == val2 { // если элементы равны добавляем в массив пересечения.
				inter = append(inter, val1)
			}
		}
	}

	return inter
}
