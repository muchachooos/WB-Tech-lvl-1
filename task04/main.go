package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main() {
	// Определяем количество воркеров.
	fmt.Println("Введите колличество воркеров: ")
	var countWorkers int
	fmt.Scan(&countWorkers)

	// Создаём канал случайной длины.
	in := make(chan interface{}, 78)

	// Запускаем горутину, которая записывает произвольные данные в канал.
	go func() {
		in <- []int{1, 4, 56}
		in <- "SomeData421"
		in <- 2343
		in <- "SomeData2"
		in <- "SomeData435"
		in <- "SomeData1"
		in <- 4322
		in <- "SomeData123"
	}()

	// Запускаем воркеры.
	RunWorkers(in, countWorkers)

	// Создаем канал для получения сигналов завершения программы.
	signalChan := make(chan os.Signal, 1)
	// Регистрируем канал на получение сигнала Interrupt (Ctrl+C).
	signal.Notify(signalChan, os.Interrupt)
	// Ждём сигнал.
	<-signalChan

	fmt.Println("--- Конец :) ---")
}

// RunWorkers запускает указанное количество воркеров и передает им задачи из канала in.
func RunWorkers(in chan interface{}, countWorkers int) {
	// Создаем WaitGroup
	wg := sync.WaitGroup{}

	// Запускаем воркеры в кол-ве countWorkers.
	for i := 0; i < countWorkers; i++ {
		wg.Add(1)     // Увеличиваем счетчик горутин в WaitGroup.
		go worker(in) // Передаём воркеру канал с задачами.
		wg.Done()     // Уменьшаем счетчик горутин в WaitGroup.
	}
	// Ждём пока все горутины запустятся.
	wg.Wait()
}

// worker выполняет задачи из канала in.
func worker(in chan interface{}) {
	for task := range in {
		// Выводим данные в stdout.
		fmt.Println(task)
	}
	// Цикл завершиться когда канал будет закрыт - тогда воркер закончит свою работу, иначе цикл ждёт.
}
