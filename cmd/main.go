package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Создаем канал для передачи данных
	dataChan := make(chan int)

	// Запускаем горутину, которая записывает данные в канал
	go func() {
		for i := 0; ; i++ {
			dataChan <- i
		}
	}()

	// Определяем количество воркеров
	numWorkers := 5

	// Создаем WaitGroup для ожидания завершения всех воркеров
	wg := sync.WaitGroup{}
	wg.Add(numWorkers)

	// Запускаем воркеры
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			for data := range dataChan {
				fmt.Println(data)
			}
		}()
	}

	// Ждем сигнала завершения программы
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Закрываем канал и ждем завершения всех воркеров
	close(dataChan)
	wg.Wait()
	fmt.Println("All workers finished")
}
