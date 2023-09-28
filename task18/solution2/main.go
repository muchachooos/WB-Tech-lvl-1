package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

type counter struct {
	val int
	mx  sync.Mutex
}

func main() {
	// Создаем WaitGroup и канал с буфером 3.
	wg := sync.WaitGroup{}
	ch := make(chan int, 3)

	// Запускаем горутину для чтения из канала.
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Выводим значения полученные из канала.
		fmt.Println("ints:", reader(ch))

	}()

	// Создаём счётчик.
	cnt := counter{}

	// Устанавливаем колл-во задач.
	countTasks := 33

	// Запускаем отправку сообщений в канал.
	sender(countTasks, ch, &cnt)

	// Закрываем канал, как только все сообщения будут отправленны.
	close(ch)
	wg.Wait()

	// Выводим значение счётчика.
	fmt.Println("cnt value:", cnt.val)
}

// sender отправляет сообщение в канал.
func sender(countTask int, ch chan<- int, cnt *counter) {
	wg := sync.WaitGroup{}

	// Запускаем горутины для отправки сообщений в канал.
	for i := 0; i < countTask; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- rand.Intn(99) // Отправляем значение в канал.
			cnt.increment()     // Инкрементируем счётчик.
			fmt.Println(cnt.val)
		}()
		time.Sleep(time.Second / 2)
	}

	wg.Wait()
}

// reader читает из канала и возвращает массив с полученными значениями.
func reader(ch <-chan int) []int {
	ints := make([]int, 0)

	// Читаем сообщения из канала до тех пор, пока он не закроется.
	// и добавляем их в слайс.
	for msg := range ch {
		ints = append(ints, msg)
	}

	// Возвращаем результат.
	return ints
}

// increment увеличивает значение счетчика на 1.
func (c *counter) increment() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.val++
}

// decrement уменьшает значение счетчика на 1
func (c *counter) decrement() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.val--
}
