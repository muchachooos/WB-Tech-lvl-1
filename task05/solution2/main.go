package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Создание канала для передачи данных.
	dataChan := make(chan int)
	// Создаём WaitGroup.
	wg := sync.WaitGroup{}

	// Запуск горутины для отправки данных в канал.
	go func() {
		// Отправляем в канал случайное число каждую секунду.
		for {
			dataChan <- rand.Int()
			time.Sleep(time.Second)
		}
	}()

	// Увеличиваем счетчик горутин в WaitGroup.
	wg.Add(1)

	// Запускаем горутину для получения данных из канала и их обработки.
	go func() {
		waitChan := time.After(5 * time.Second) // Создаём канал в который поступит сообщение через 5 секунд.
		for {
			// Ожидаем данные из каналов waitChan и dataChan.
			select {
			case <-waitChan: // Если в waitChan есть данные (то есть прошло 5 секунд), то завершаем ф-цию.
				fmt.Println("Timeout")
				wg.Done()
				return
			case data := <-dataChan: // Если в dataChan есть данные, то выводим их в консоль.
				fmt.Println("Received:", data)
			}
		}
	}()

	// Ждём пока горутина завершит свою работу.
	wg.Wait()
	fmt.Println("End")
}
