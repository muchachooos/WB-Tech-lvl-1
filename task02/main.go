package main

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

import (
	"fmt"
	"sync"
)

func main() {
	// Создаём массив.
	arr := []int{2, 4, 6, 8, 10}

	// Создаём WaitGroup.
	var wg sync.WaitGroup

	// Для каждого элемента запускаем новую горутину которая выполняет расчет.
	for _, num := range arr {
		wg.Add(1) // Увеличиваем счетчик горутин в WaitGroup.
		go func(n int) {
			square := n * n
			fmt.Println(square)
			wg.Done() // Уменьшаем счетчик горутин в WaitGroup.
		}(num)
	}

	// Ждём пока все горутины завершат свою работу.
	wg.Wait()
}