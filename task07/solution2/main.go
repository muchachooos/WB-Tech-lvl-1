package main

import (
	"math/rand"
	"sync"
)

// Реализовать конкурентную запись данных в map.

func main() {
	// Создаём Map.
	data := make(map[int]int)
	// Создаём WaitGroup.
	wg := sync.WaitGroup{}
	// Создаём Mutex.
	mx := sync.Mutex{}

	// Запускаем 10 горутин, которые записывают случайное значение в map.
	for i := 0; i < 10; i++ {
		wg.Add(1) // Увеличиваем счетчик горутин в WaitGroup.
		go func() {
			mx.Lock() // Блокируем Mutex.

			data[rand.Intn(100)] = rand.Intn(100) // Устанавливаем значения в data.

			mx.Unlock() // Разблокируем Mutex.
			wg.Done()   // Уменьшаем счетчик горутин в WaitGroup.
		}()
	}

	// Ждём пока все горутины завершат работу.
	wg.Wait()
}
