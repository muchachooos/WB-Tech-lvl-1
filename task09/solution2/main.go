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
	// Создаем каналы для передачи данных между горутинами.
	in := make(chan int)
	out := make(chan int)

	// Запускаем горутину для чтения данных из массива data и отправки их в канал in.
	go reader(data, in)
	// Запускаем горутину для умножения чисел из канала in на 2 и отправки результатов в канал out.
	go multiplier(in, out)

	wg := sync.WaitGroup{}

	// Запускаем горутину которая выводит в stdout все числа из канала out.
	wg.Add(1)
	go printRes(out, &wg)

	wg.Wait()
}

func reader(data []int, in chan int) {
	// Отправляем в канал in числа из массива.
	for _, val := range data {
		in <- val
	}

	// Закрываем канал, когда все числа отправлены.
	close(in)
}

func multiplier(in, out chan int) {
	// Читаем значения из канала in, умножаем и отправляем в канал out.
	for val := range in {
		out <- val * 2
	}

	// Закрываем канал, когда все числа отправлены.
	close(out)
}

func printRes(out chan int, wg *sync.WaitGroup) {
	// Читаем значения из канала out и выводим в stdout.
	for num := range out {
		fmt.Println(num)
	}

	wg.Done()
}
