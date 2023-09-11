package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Реализовать конкурентную запись данных в map.

// Определяем структуру которая содержит Map и Mutex.
type threadSafetyMap struct {
	data map[int]int
	mx   sync.Mutex
}

func main() {
	// Создаём экземпляр threadSafetyMap и создаём в нём пустую map.
	m := &threadSafetyMap{
		data: make(map[int]int),
	}

	// Создаём WaitGroup.
	wg := sync.WaitGroup{}

	// Запускаем 10 горутин, которые вызывают метод set.
	for i := 0; i < 10; i++ {
		wg.Add(1) // Увеличиваем счетчик горутин в WaitGroup.
		go func() {
			// Вызываем метод set и передаём в него случайные значения.
			m.set(rand.Intn(100), rand.Intn(100))
			// Уменьшаем счетчик горутин в WaitGroup.
			wg.Done()
		}()
	}

	// Ждём пока все горутины завершат работу.
	wg.Wait()
	// Выводим значения поля data.
	fmt.Println(m.data)
}

// set устанавливает значение val по ключу key в поле data.
func (tsm *threadSafetyMap) set(key, val int) {
	// Блокируем Mutex.
	tsm.mx.Lock()
	// Разблокируем Mutex при выходе из ф-ции.
	defer tsm.mx.Unlock()

	// Устанавливаем значения в data
	tsm.data[key] = val
}
