package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {
	// Создаём массив.
	arr := []int{2, 4, 6, 8, 10}

	// Создаём WaitGroup.
	var wg sync.WaitGroup

	// Создаём Mutex.
	mx := sync.Mutex{}

	var sumSquare int

	// Для каждого элемента запускаем новую горутину, которая выполняет расчет.
	for _, num := range arr {
		wg.Add(1) // Увеличиваем счетчик горутин в WaitGroup.
		go func(n int) {
			square := n * n
			mx.Lock() // Блокируем Mutex.
			sumSquare += square
			mx.Unlock() // Разблокируем Mutex после завершения работы с общими данными.
			wg.Done()   // Уменьшаем счетчик горутин в WaitGroup.
		}(num)
	}

	// Ждём пока все горутины завершат свою работу.
	wg.Wait()
	fmt.Println("Сумма квадратов:", sumSquare)
}
