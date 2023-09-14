package main

import (
	"fmt"
	"sync"
)

// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

func main() {
	// Создаём массив с данными.
	data := []int{3, 33, 133, 233, 333, 433, 533, 633, 733}

	in := reader(data)
	out := multiplier(in)

	wg := sync.WaitGroup{}

	// Запускаем горутину которая выводит в stdout все числа из канала out.
	wg.Add(1)
	go func() {
		for num := range out {
			fmt.Println(num)
		}
		wg.Done()
	}()

	wg.Wait()
}

// reader читает числа из массива, отправляет их в канал и возвращает его.
func reader(data []int) chan int {
	// Создаём канал.
	in := make(chan int)

	// Запускаем горутину для отправки чисел из массива data в канал in.
	go func() {
		for _, val := range data {
			in <- val
		}
		close(in) // Закрываем канал после отправки всех чисел.
	}()

	// Возращаем канал.
	return in
}

// multiplier читает числа из канала in, умножает их, отправляет в канал out и возвращает его.
func multiplier(in chan int) chan int {
	// Создаём канал для результатов.
	out := make(chan int)

	// Запускам горутину для умножения чисел из канала in и отправки в канал out.
	go func() {
		for val := range in {
			out <- val * 2
		}
		close(out) // Закрываем канал после отправки всех чисел.
	}()

	// Возвращаем канал.
	return out
}
