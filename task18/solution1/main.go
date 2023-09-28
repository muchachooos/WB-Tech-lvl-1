package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

type counter struct {
	val int
	mx  sync.Mutex
}

func main() {
	// Создаем WaitGroup и счётчик.
	wg := sync.WaitGroup{}
	counter := counter{}

	// Устанавливаем колл-во задач.
	countTasks := 10

	// Запускаем горутины которые инкрементируют счётчик.
	for i := 0; i < countTasks; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.increment()
		}()
	}

	wg.Wait()
	fmt.Println(counter.val)
}

// increment увеличивает значение счетчика на 1.
func (c *counter) increment() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.val++
}
