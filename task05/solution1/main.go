package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения
// в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

func main() {
	// Создаём канал.
	in := make(chan int, 1)

	// Отправляем данные в канал.
	go sendToChannel(in)
	// Зачитываем данные из канала.
	go readFromChanel(in)

	// Ждём 5 секунд и завершаем программу.
	time.Sleep(time.Second * 5)
}

func sendToChannel(in chan int) {
	// Бесконечно отправляем в канал случайное число.
	for {
		in <- rand.Int()
		fmt.Println("Send")
		time.Sleep(time.Second)
	}
}

func readFromChanel(in chan int) {
	// Считываем данные из канала.
	for i := range in {
		fmt.Println("Received: ", i)
	}
}
